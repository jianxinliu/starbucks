create table material
(
    id          bigint unsigned auto_increment,

    material_id varchar(30) not null comment '原料 id',
    meta_id     varchar(30) not null comment '原料类型 ID',


    primary key (id),
    unique key (material_id)

) ENGINE = InnoDB comment '原料表'