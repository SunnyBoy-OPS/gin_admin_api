-- auto-generated definition 用户表
create table `sys.UserInfo`
(
    Id           int unsigned auto_increment primary key comment '主键',
    UserName     varchar(20) unique not null comment '账号',
    UserPassword varchar(20)        not null comment '密码',
    NickNane     varchar(20)        not null comment '姓名',
    Icon         varchar(40)                 default null comment '头像',
    Sex          int                not null default '1' comment '性别(1->男,n->女)',
    Email        varchar(30)                 default null comment '邮箱',
    Status       int                not null default '1' comment '账号启用状态(1->启用,n->停止)',
    DepartmentId int                not null UNIQUE comment '部门ID',
    CreateTime   datetime                    default CURRENT_TIMESTAMP not null comment '创建时间'
)ENGINE = InnoDB default CHARSET = UTF8 ROW_FORMAT = DYNAMIC COMMENT '用户信息表';

