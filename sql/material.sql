create table material
(
    id          bigint unsigned auto_increment,
    created_at  datetime(3) null,
    updated_at  datetime(3) null,
    deleted_at  datetime(3) null,

    material_id varchar(30) not null comment '原料 id',
    meta_id     varchar(30) not null comment '原料类型 ID',


    primary key (id),
    unique key (material_id)

) ENGINE = InnoDB comment '原料表'