create table wallet
(
    id         bigint unsigned auto_increment,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,

    user_id    varchar(30) not null,
    balance    bigint default 0 comment '用户余额，单位： 分',

    primary key (id),
    unique key (user_id)

) ENGINE = InnoDB comment '用户余额表'