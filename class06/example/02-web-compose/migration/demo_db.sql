-- auto-generated definition
create schema s_demo;

alter schema s_demo owner to postgres;

create table s_demo.t_user
(
    id       bigserial
        constraint t_user_pk
            primary key,
    username varchar(128) default ''::character varying not null,
    password varchar(512) default ''::character varying not null
);

alter table s_demo.t_user
    owner to postgres;

create unique index t_user_username_uindex
    on s_demo.t_user (username);


