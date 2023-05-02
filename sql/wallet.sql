create table wallet
(
    id      bigint unsigned auto_increment,

    user_id varchar(30) not null,
    balance int default 0 comment '用户余额，单位： 分',

    primary key (id),
    unique key (user_id)

) ENGINE = InnoDB comment '用户余额表'