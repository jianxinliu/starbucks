create table recipe_item
(
    id          bigint unsigned auto_increment,

    recipe_id   varchar(30) not null,
    material_id varchar(30) not null comment '使用那种原料',
    quantity    numeric default 1 comment '原料数量',

    primary key (id),
    unique key (recipe_id, material_id)

) ENGINE = InnoDB comment '配方项详细表'