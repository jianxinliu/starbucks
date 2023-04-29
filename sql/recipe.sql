create table recipe
(
    id          bigint unsigned auto_increment,
    created_at  datetime(3) null,
    updated_at  datetime(3) null,
    deleted_at  datetime(3) null,

    recipe_id   varchar(30) not null,
    recipe_name varchar(30) not null comment '配方名称',

    primary key (id),
    unique key (recipe_id)
) ENGINE = InnoDB comment '配方表'