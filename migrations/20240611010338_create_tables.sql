-- +goose Up


create table if not exists "order"
(
    order_uid          varchar(255) unique      not null,
    track_number       varchar                  not null,
    entry              varchar                  not null,
    locale             varchar(10)              not null,
    internal_signature varchar(255)             not null,
    customer_id        varchar(100)             not null,
    delivery_service   varchar(255)             not null,
    shardkey           varchar(255)             not null,
    sm_id              integer                  not null,
    date_created       timestamp with time zone not null,
    oof_shard          varchar(255)             not null
);

create table if not exists "delivery"
(
    order_uid varchar(255) not null,
    name      varchar(255) not null,
    phone     varchar(255) not null,
    zip       integer      not null,
    city      varchar(255) not null,
    address   varchar(255) not null,
    region    varchar(255) not null,
    email     varchar(255) not null,

    constraint ch_phone
        check (phone ~* '^\+7\d{10}$'),
    constraint ch_email
        check (email ~* '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$')
);

create table if not exists "payment"
(
    transaction   varchar(255)        not null,
    request_id    varchar(255) unique not null,
    currency      varchar(255)        not null,
    provider      varchar(255)        not null,
    amount        integer             not null,
    payment_dt    integer             not null,
    bank          varchar(255)        not null,
    delivery_cost numeric(10, 2)      not null,
    goods_total   integer             not null,
    custom_fee    numeric(10, 2)      not null
);

create table if not exists "items"
(
    chrt_id      bigint primary key                          not null,
    track_number varchar(255)                                not null,
    price        numeric(10, 2)                              not null,
    rid          varchar(255) unique                         not null,
    name         varchar(255)                                not null,
    sale         integer check ( sale >= 0 and sale <= 100 ) not null,
    size         varchar(255)                                not null,
    total_price  numeric(10, 2)                              not null,
    nm_id        integer                                     not null,
    brand        varchar(255)                                not null,
    status       integer                                     not null
);

-- +goose Down
drop table if exists "order" cascade;
drop table if exists delivery cascade;
drop table if exists payment cascade;
drop table if exists items cascade;
