create database coupon
with owner postgres;


create table consume
(
  id         varchar(36) not null
    constraint consume_pkey
    primary key,
  cuscopid   varchar(36) not null,
  cust_phone varchar(36) not null,
  mch_phone  varchar(36) not null,
  amount     integer,
  success_at timestamp   not null
);

alter table consume
  owner to postgres;

create table coupon
(
  id         varchar(36)       not null
    constraint coupon_pkey
    primary key,
  title      varchar(20)       not null,
  code       varchar(6)        not null,
  amount     integer default 0 not null,
  start_time timestamp         not null,
  exp_time   timestamp         not null,
  content    varchar(30),
  remarks    varchar(30)
);

alter table coupon
  owner to postgres;

create table cus_cop
(
  copid  varchar(36)       not null
    constraint cus_cop_pk
    primary key,
  phone  varchar(20)       not null,
  status integer default 0 not null
);

comment on column cus_cop.status
is '0:待核销，1：以核销，2：以过期';

alter table cus_cop
  owner to postgres;

create table customer
(
  phone varchar(20)       not null
    constraint customer_pkey
    primary key,
  name  varchar(15),
  type  integer default 1 not null
);

comment on column customer.type
is '1:顾客，2:商户';

alter table customer
  owner to postgres;

