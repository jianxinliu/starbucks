create table product_group
(
    id         bigint unsigned auto_increment,

    group_id   varchar(30) not null,
    group_name varchar(50) comment '分类名称: 咖啡，果汁……',
    group_desc varchar(300),

    primary key (id),
    unique key (group_id)
) ENGINE = InnoDB comment '产品分类表';