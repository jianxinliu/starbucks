create table user
(
    id         bigint unsigned auto_increment,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,

    user_id    varchar(30) not null,
    user_name  varchar(30) not null,
    password   varchar(50) not null,
    user_type  varchar(20) comment '用户类型 normal, vip……',

    primary key (id),
    unique key (user_id)
) ENGINE = InnoDB