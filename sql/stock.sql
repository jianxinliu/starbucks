create table stock
(
    id          bigint unsigned auto_increment,

    stock_id    varchar(30)         not null comment '库存 id',
    material_id varchar(30)         not null comment '存储的哪种原料',

    amount      int       default 0 not null comment '入库数量',
    `count`     int       default 0 not null comment '当前数量',

    in_time     timestamp default now() comment '原料进库时间',

    primary key (id),
    unique key (stock_id)

) ENGINE = InnoDB comment '原料表'