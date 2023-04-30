create table products
(
    id          bigint unsigned auto_increment,

    product_id  varchar(130) not null,
    name        varchar(100) not null,
    description varchar(100),
    image       varchar(500),

    group_id    varchar(30)  not null comment '产品分组',
    price       int comment '价格，单位：分',
    discount    numeric comment '折扣',

    primary key (id),
    unique key (product_id)
) ENGINE = InnoDB;