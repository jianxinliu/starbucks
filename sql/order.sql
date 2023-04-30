create table `order`
(
    id          bigint unsigned auto_increment,

    order_id    varchar(30) not null,
    product_id  varchar(30) not null comment '哪个产品的订单, o-cafe-xxx, o-vip-xxx, ……',
    status      int       default 0 comment '订单状态。0,1,2,……',
    create_time timestamp default current_timestamp comment '订单创建时间',
    finish_time timestamp comment '订单结束时间',

    primary key (id),
    unique key (order_id)
) ENGINE = InnoDB