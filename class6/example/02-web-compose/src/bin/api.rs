use crate::demo_service::{LoginReq, RegisterReq};
use axum::extract::{Form, State};
use axum::routing::post;
use axum::{BoxError, Json, Router, Server};
use demo_service::demo_srv_client::DemoSrvClient;
use serde::{Deserialize, Serialize};
use tonic::transport::Channel;
use tower_http::trace::{DefaultOnRequest, DefaultOnResponse};
use tower_http::LatencyUnit;
use tracing::Level;

pub mod demo_service {
    tonic::include_proto!("demo");
}

#[derive(Deserialize)]
pub struct SignReq {
    username: String,
    password: String,
}

#[derive(Serialize)]
pub struct SignInRes {
    code: i32,
    msg: String,
    token: String,
}

#[derive(Serialize)]
pub struct SignUpRes {
    code: i32,
    msg: String,
}

impl Into<SignInRes> for demo_service::LoginRes {
    fn into(self) -> SignInRes {
        SignInRes {
            code: self.code,
            msg: self.msg,
            token: self.token,
        }
    }
}

impl Into<SignUpRes> for demo_service::RegisterRes {
    fn into(self) -> SignUpRes {
        SignUpRes {
            code: self.code,
            msg: self.msg,
        }
    }
}

#[derive(Clone, Debug)]
struct AppContext<T> {
    rpc_conn: DemoSrvClient<T>,
}

#[tokio::main]
async fn main() -> Result<(), BoxError> {
    tracing_subscriber::fmt::init();

    let port_api = std::env::var("API_PORT").unwrap_or("3000".into());

    let host_rpc = std::env::var("RPC_HOST").unwrap_or("127.0.0.1".into());
    let port_rpc = std::env::var("RPC_PORT").unwrap_or("50051".into());
    let client = DemoSrvClient::connect(format!("http://{}:{}", host_rpc, port_rpc)).await?;

    let ctx = AppContext { rpc_conn: client };
    let app = Router::new()
        .route("/login", post(handle_login))
        .route("/register", post(handle_register))
        .layer(
            tower_http::trace::TraceLayer::new_for_http()
                .on_request(DefaultOnRequest::new().level(Level::INFO))
                .on_response(
                    DefaultOnResponse::new()
                        .level(Level::INFO)
                        .latency_unit(LatencyUnit::Micros),
                ),
        )
        .with_state(ctx);

    println!("DemoService API listening on port {}", port_api);

    Server::bind(&format!("0.0.0.0:{}", port_api).parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
    Ok(())
}

async fn handle_register(
    State(AppContext { mut rpc_conn }): State<AppContext<Channel>>,
    Form(SignReq { username, password }): Form<SignReq>,
) -> Json<SignUpRes> {
    let resp = rpc_conn.register(RegisterReq { username, password }).await;
    match resp.map(tonic::Response::into_inner) {
        Ok(res) => Json(res.into()),
        Err(_) => Json(SignUpRes {
            code: 500,
            msg: "internal error".into(),
        }),
    }
}

async fn handle_login(
    State(AppContext { mut rpc_conn }): State<AppContext<Channel>>,
    Form(SignReq { username, password }): Form<SignReq>,
) -> Json<SignInRes> {
    let resp = rpc_conn.login(LoginReq { username, password }).await;

    return match resp.map(tonic::Response::into_inner) {
        Ok(res) => Json(res.into()),
        Err(_) => Json(SignInRes {
            code: 500,
            msg: "internal error".into(),
            token: "".into(),
        }),
    };
}
