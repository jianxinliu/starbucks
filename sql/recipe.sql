create table recipe
(
    id          bigint unsigned auto_increment,

    recipe_id   varchar(30) not null,
    recipe_name varchar(30) not null comment '配方名称',
    description varchar(300) comment '配方说明',
    `version`   int         not null default 0 comment '配方版本，变更时升级版本',

    primary key (id),
    unique key (recipe_id, `version`)
) ENGINE = InnoDB comment '配方表'