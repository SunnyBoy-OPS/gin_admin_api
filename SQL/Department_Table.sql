-- 部门表Department
create table `Department`(
                             DepartmentId int unsigned auto_increment primary key comment '部门ID',
                             DepartmentName varchar(10) not null unique comment '部门名称',
                             CreateTime   datetime default CURRENT_TIMESTAMP not null comment '创建时间'
)ENGINE = InnoDB default CHARSET = UTF8 ROW_FORMAT = DYNAMIC COMMENT '部门表';