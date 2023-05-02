create table payment
(
    id           bigint unsigned auto_increment,
    pay_id       varchar(30) not null,
    user_id      varchar(30) not null,
    order_id     varchar(30) not null,
    payed_amount int         not null comment '已支付金额，单位：分',
    payed_time   timestamp default current_timestamp comment '支付时间',

    primary key (id),
    unique key (pay_id)
) ENGINE = InnoDB comment '支付记录表'