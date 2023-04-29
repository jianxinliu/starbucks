create table material_meta
(
    id            bigint unsigned auto_increment,
    created_at    datetime(3) null,
    updated_at    datetime(3) null,
    deleted_at    datetime(3) null,

    material_id   varchar(30) not null comment '原料 ID',
    material_type varchar(30) not null comment '原料类型。豆子，糖，奶，巧克力……',
    name          varchar(50),
    unit          varchar(30) not null comment '单位',

    primary key (id),
    unique key (material_id)

) ENGINE = InnoDB comment '原料类型表'