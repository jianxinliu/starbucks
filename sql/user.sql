create table user
(
    id             bigint unsigned auto_increment,

    user_id        varchar(30) not null,
    user_name      varchar(30) not null,
    password       varchar(50) not null,
    user_type      varchar(20) default 'normal' comment '用户类型 normal, vip1, vip2……',

    vip_start_time timestamp comment '会员起始时间',
    vip_end_time   timestamp comment '会员到期时间',

    primary key (id),
    unique key (user_id)
) ENGINE = InnoDB