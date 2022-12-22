use axum::BoxError;
use chrono::Utc;
use demo_service::demo_srv_server::{DemoSrv, DemoSrvServer};
use demo_service::{LoginReq, LoginRes, RegisterReq, RegisterRes};
use redis::{AsyncCommands, Client};
use sha256::digest;
use sqlx::postgres::{PgConnectOptions, PgPoolOptions};
use sqlx::{ConnectOptions, Error, Pool, Postgres};
use std::net::SocketAddr;
use tonic::transport::Server;
use tonic::{Request, Response, Status};
use tracing::{event, Level};

pub mod demo_service {
    tonic::include_proto!("demo");
}

pub struct DemoService {
    pool: Pool<Postgres>,
    salt: String,
    redis: Client,
}

#[tonic::async_trait]
impl DemoSrv for DemoService {
    async fn login(&self, req: Request<LoginReq>) -> Result<Response<LoginRes>, Status> {
        let req = req.into_inner();
        let result = sqlx::query_as::<_, (String,)>(
            "SELECT password FROM s_demo.t_user WHERE username = $1 LIMIT 1",
        )
        .bind(&req.username)
        .fetch_one(&self.pool)
        .await;
        return match result {
            Ok((pwd,)) => {
                let hashed = digest(format!("{}{}", &req.password, &self.salt));
                if pwd == hashed {
                    let token = format!(
                        "{}.{}.{}",
                        &req.username,
                        digest(format!("{}{}", &req.username, &self.salt)),
                        Utc::now().timestamp()
                    );
                    let mut conn = self.redis.get_async_connection().await
                        .map_err(|e| { Status::from_error(Box::new(e)) })?;
                    let _: () = conn.set(format!("user:registered:{}", &req.username), "").await
                        .map_err(|e| { Status::from_error(Box::new(e)) })?;
                    Ok(Response::new(LoginRes {
                        code: 200,
                        msg: "ok".to_string(),
                        token,
                    }))
                } else {
                    Ok(Response::new(LoginRes {
                        code: 400,
                        msg: "invalid password".to_string(),
                        token: String::new(),
                    }))
                }
            }
            Err(err) => match err {
                Error::RowNotFound => Ok(Response::new(LoginRes {
                    code: 400,
                    msg: "no user".to_string(),
                    token: "".to_string(),
                })),
                _ => {
                    event!(
                        Level::ERROR,
                        target = "demo-rpc",
                        "error happened when connect to database, err: {:?}",
                        err
                    );
                    Err(Status::from_error(Box::new(err)))
                }
            },
        };
    }

    async fn register(&self, req: Request<RegisterReq>) -> Result<Response<RegisterRes>, Status> {
        let req = req.into_inner();
        let mut conn = self.redis.get_async_connection().await
            .map_err(|e| { Status::from_error(Box::new(e)) })?;
        let ok: bool = conn.exists(format!("user:registered:{}", &req.username)).await
            .map_err(|e| { Status::from_error(Box::new(e)) })?;

        if ok {
            return Ok(Response::new(RegisterRes {
                code: 400,
                msg: "user already registered".to_string()
            }))
        }

        let mut tx = self.pool.begin().await
            .map_err(|e| { Status::from_error(Box::new(e)) })?;

        let (result,) = sqlx::query_as::<_, (i64,)>("SELECT count(*) FROM s_demo.t_user WHERE username = $1")
            .bind(&req.username)
            .fetch_one(&mut tx)
            .await.map_err(|e| {
                Status::from_error(Box::new(e))
            })?;
        if result > 0 {
            return Ok(Response::new(RegisterRes {
                code: 400,
                msg: "user already registered".to_string()
            }))
        }

        let hashed = digest(format!("{}{}", &req.password, &self.salt));
        let result = sqlx::query("INSERT INTO s_demo.t_user (username, password) VALUES ($1, $2)")
            .bind(&req.username)
            .bind(hashed)
            .execute(&mut tx)
            .await;

        tx.commit().await.map_err(|e| {
            Status::from_error(Box::new(e))
        })?;

        let _: () = conn.set(format!("user:registered:{}", &req.username), "").await
            .map_err(|e| { Status::from_error(Box::new(e)) })?;

        result
            .and_then(|_| {
                Ok(Response::new(RegisterRes {
                    code: 200,
                    msg: "register success".to_string(),
                }))
            })
            .map_err(|err| {
                event!(
                    Level::ERROR,
                    target = "demo-rpc",
                    "error happened when register user {}, err: {:?}",
                    &req.username,
                    err
                );
                Status::from_error(Box::new(err))
            })
    }
}

#[tokio::main]
async fn main() -> Result<(), BoxError> {
    tracing_subscriber::fmt::init();
    let salt = std::env::var("PWD_SALT").unwrap_or("114514".into());
    let port_rpc = std::env::var("RPC_PORT").unwrap_or("50051".into());
    let addr: SocketAddr = format!("0.0.0.0:{}", port_rpc).parse().unwrap();
    let db_host = std::env::var("DB_HOST").unwrap_or("localhost".into());
    let db_port: u16 = std::env::var("DB_PORT").unwrap_or("5432".into()).parse().unwrap();
    let db_username = std::env::var("DB_USERNAME").unwrap_or("postgres".into());
    let db_pwd = std::env::var("DB_PWD").unwrap_or("deepdarkfancy".into());
    let db_name = std::env::var("DB_NAME").unwrap_or("demo_db".into());
    let redis_host = std::env::var("REDIS_HOST").unwrap_or("localhost".into());
    let redis_port = std::env::var("REDIS_PORT").unwrap_or("6379".into());

    let mut opt = PgConnectOptions::new()
        .host(&*db_host)
        .port(db_port)
        .username(&*db_username)
        .password(&*db_pwd)
        .database(&*db_name);
    opt.disable_statement_logging();

    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect_with(opt)
        .await?;

    let redis = Client::open(format!("redis://{}:{}/", redis_host, redis_port)).unwrap();

    let srv = DemoService { pool, salt, redis };

    println!("DemoService listening on {}", addr);

    Server::builder()
        .add_service(DemoSrvServer::new(srv))
        .serve(addr)
        .await?;

    Ok(())
}
