-- -------------------------------------------------------------
-- TablePlus 5.3.5(494)
--
-- https://tableplus.com/
--
-- Database: fkratos_sys
-- Generation Time: 2023-05-24 20:47:31.6130
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."sys_admin";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_admin" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "tenant_id" uuid NOT NULL,
    "username" varchar(50) NOT NULL,
    "password" varchar(128) NOT NULL,
    "nickname" varchar(50) NOT NULL,
    "avatar" varchar(255),
    "gender" int2 NOT NULL DEFAULT 0,
    "email" varchar(50),
    "mobile" varchar(15),
    "job_id" uuid,
    "dept_id" uuid,
    "role_ids" json,
    "salt" varchar(32) NOT NULL,
    "status" int2 NOT NULL DEFAULT 1,
    "motto" varchar(255),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_admin"."id" IS '编号';
COMMENT ON COLUMN "public"."sys_admin"."tenant_id" IS '租户ID';
COMMENT ON COLUMN "public"."sys_admin"."username" IS '用户名';
COMMENT ON COLUMN "public"."sys_admin"."password" IS '密码';
COMMENT ON COLUMN "public"."sys_admin"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."sys_admin"."avatar" IS '头像';
COMMENT ON COLUMN "public"."sys_admin"."gender" IS '0=保密 1=女 2=男';
COMMENT ON COLUMN "public"."sys_admin"."email" IS '邮件';
COMMENT ON COLUMN "public"."sys_admin"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."sys_admin"."job_id" IS '岗位';
COMMENT ON COLUMN "public"."sys_admin"."dept_id" IS '部门';
COMMENT ON COLUMN "public"."sys_admin"."role_ids" IS '角色集';
COMMENT ON COLUMN "public"."sys_admin"."salt" IS '盐值';
COMMENT ON COLUMN "public"."sys_admin"."status" IS '0=禁用 1=开启';
COMMENT ON COLUMN "public"."sys_admin"."motto" IS '个性签名';
COMMENT ON COLUMN "public"."sys_admin"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."sys_admin"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."sys_admin"."deleted_at" IS '删除时间';

DROP TABLE IF EXISTS "public"."sys_api";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_api" (
    "id" uuid NOT NULL,
    "tenant_id" uuid,
    "permission_id" uuid NOT NULL,
    "method" varchar(32) NOT NULL,
    "path" varchar(255) NOT NULL,
    "desc" varchar(255) NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_api"."id" IS '编号';
COMMENT ON COLUMN "public"."sys_api"."tenant_id" IS '租户ID';
COMMENT ON COLUMN "public"."sys_api"."permission_id" IS '权限Id';
COMMENT ON COLUMN "public"."sys_api"."method" IS '方法';
COMMENT ON COLUMN "public"."sys_api"."path" IS '路径';
COMMENT ON COLUMN "public"."sys_api"."desc" IS '描述';
COMMENT ON COLUMN "public"."sys_api"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."sys_api"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."sys_api"."deleted_at" IS '删除时间';

DROP TABLE IF EXISTS "public"."sys_dept";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_dept" (
    "id" uuid NOT NULL,
    "tenant_id" uuid,
    "pid" uuid NOT NULL,
    "name" varchar(50) NOT NULL,
    "full_name" varchar(50) NOT NULL,
    "responsible" varchar(20),
    "phone" varchar(20),
    "email" varchar(255),
    "type" int2 NOT NULL,
    "status" int2 NOT NULL,
    "sort" int8 NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_dept"."id" IS '编号';
COMMENT ON COLUMN "public"."sys_dept"."tenant_id" IS '租户ID';
COMMENT ON COLUMN "public"."sys_dept"."pid" IS '父级id';
COMMENT ON COLUMN "public"."sys_dept"."name" IS '部门简称';
COMMENT ON COLUMN "public"."sys_dept"."full_name" IS '部门全称';
COMMENT ON COLUMN "public"."sys_dept"."responsible" IS '负责人';
COMMENT ON COLUMN "public"."sys_dept"."phone" IS '负责人电话';
COMMENT ON COLUMN "public"."sys_dept"."email" IS '负责人邮箱';
COMMENT ON COLUMN "public"."sys_dept"."type" IS '1=公司 2=子公司 3=部门';
COMMENT ON COLUMN "public"."sys_dept"."status" IS '0=禁用 1=开启';
COMMENT ON COLUMN "public"."sys_dept"."sort" IS '排序值';
COMMENT ON COLUMN "public"."sys_dept"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."sys_dept"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."sys_dept"."deleted_at" IS '删除时间';

DROP TABLE IF EXISTS "public"."sys_dict";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_dict" (
    "id" uuid NOT NULL,
    "tenant_id" uuid,
    "pid" uuid NOT NULL,
    "name" varchar(50) NOT NULL,
    "type" int2 NOT NULL,
    "unique_key" varchar(50) NOT NULL,
    "value" varchar(2048) NOT NULL,
    "status" int2 NOT NULL,
    "sort" numeric(20,0) NOT NULL,
    "remark" varchar(200) NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_dict"."id" IS '编号';
COMMENT ON COLUMN "public"."sys_dict"."tenant_id" IS '租户ID';
COMMENT ON COLUMN "public"."sys_dict"."pid" IS '0=配置集 !0=父级id';
COMMENT ON COLUMN "public"."sys_dict"."name" IS '名称';
COMMENT ON COLUMN "public"."sys_dict"."type" IS '1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单图 10多图 11单文件 12多文件';
COMMENT ON COLUMN "public"."sys_dict"."unique_key" IS '唯一值';
COMMENT ON COLUMN "public"."sys_dict"."value" IS '配置值';
COMMENT ON COLUMN "public"."sys_dict"."status" IS '0=禁用 1=开启';
COMMENT ON COLUMN "public"."sys_dict"."sort" IS '排序值';
COMMENT ON COLUMN "public"."sys_dict"."remark" IS '备注';
COMMENT ON COLUMN "public"."sys_dict"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."sys_dict"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."sys_dict"."deleted_at" IS '删除时间';

DROP TABLE IF EXISTS "public"."sys_job";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_job" (
    "id" uuid NOT NULL,
    "tenant_id" uuid,
    "name" varchar(50) NOT NULL,
    "code" varchar(32),
    "remark" varchar(255),
    "sort" int8 NOT NULL,
    "status" int2 NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_job"."id" IS '编号';
COMMENT ON COLUMN "public"."sys_job"."tenant_id" IS '租户ID';
COMMENT ON COLUMN "public"."sys_job"."name" IS '岗位名称';
COMMENT ON COLUMN "public"."sys_job"."code" IS '岗位编码';
COMMENT ON COLUMN "public"."sys_job"."remark" IS '备注';
COMMENT ON COLUMN "public"."sys_job"."sort" IS '排序值';
COMMENT ON COLUMN "public"."sys_job"."status" IS '0=禁用 1=开启 ';
COMMENT ON COLUMN "public"."sys_job"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."sys_job"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."sys_job"."deleted_at" IS '删除时间';

DROP TABLE IF EXISTS "public"."sys_login_log";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_login_log" (
    "id" uuid,
    "tenant_id" uuid,
    "admin_id" uuid,
    "ip" varchar(32),
    "useragent" varchar(255),
    "login_time" timestamptz
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_login_log"."tenant_id" IS '组织ID';
COMMENT ON COLUMN "public"."sys_login_log"."admin_id" IS '用户ID';
COMMENT ON COLUMN "public"."sys_login_log"."ip" IS 'ip';
COMMENT ON COLUMN "public"."sys_login_log"."useragent" IS '浏览器标识';
COMMENT ON COLUMN "public"."sys_login_log"."login_time" IS '登录时间';

DROP TABLE IF EXISTS "public"."sys_operation_log";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_operation_log" (
    "id" uuid NOT NULL,
    "tenant_id" uuid,
    "admin_id" uuid NOT NULL,
    "ip" varchar(32) NOT NULL,
    "uri" varchar(200) NOT NULL,
    "useragent" varchar(255),
    "header" json,
    "req" json,
    "resp" json,
    "created_at" timestamptz NOT NULL,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_operation_log"."id" IS '编号';
COMMENT ON COLUMN "public"."sys_operation_log"."tenant_id" IS '租户ID';
COMMENT ON COLUMN "public"."sys_operation_log"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."sys_operation_log"."ip" IS 'ip';
COMMENT ON COLUMN "public"."sys_operation_log"."uri" IS '请求路径';
COMMENT ON COLUMN "public"."sys_operation_log"."useragent" IS '浏览器标识';
COMMENT ON COLUMN "public"."sys_operation_log"."header" IS 'header';
COMMENT ON COLUMN "public"."sys_operation_log"."req" IS '请求数据';
COMMENT ON COLUMN "public"."sys_operation_log"."resp" IS '响应数据';
COMMENT ON COLUMN "public"."sys_operation_log"."created_at" IS '创建时间';

DROP TABLE IF EXISTS "public"."sys_permission";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_permission" (
    "id" uuid NOT NULL,
    "tenant_id" uuid,
    "pid" uuid NOT NULL,
    "type" varchar(255) NOT NULL,
    "title" varchar(50) NOT NULL,
    "name" varchar(50) NOT NULL,
    "path" varchar(100) NOT NULL,
    "icon" varchar(50) NOT NULL,
    "menu_type" varchar(255),
    "url" varchar(255) NOT NULL,
    "component" varchar(100) NOT NULL,
    "extend" varchar(255) NOT NULL,
    "remark" varchar(255) NOT NULL,
    "sort" int8 NOT NULL,
    "status" int2 NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_permission"."tenant_id" IS '租户ID';
COMMENT ON COLUMN "public"."sys_permission"."pid" IS '上级菜单';
COMMENT ON COLUMN "public"."sys_permission"."type" IS '类型:menu_dir=菜单目录,menu=菜单项,button=页面按钮';
COMMENT ON COLUMN "public"."sys_permission"."title" IS '标题';
COMMENT ON COLUMN "public"."sys_permission"."name" IS '规则名称';
COMMENT ON COLUMN "public"."sys_permission"."path" IS '路由路径';
COMMENT ON COLUMN "public"."sys_permission"."icon" IS '图标';
COMMENT ON COLUMN "public"."sys_permission"."menu_type" IS '菜单类型:tab=选项卡,link=链接,iframe=Iframe';
COMMENT ON COLUMN "public"."sys_permission"."url" IS 'Url';
COMMENT ON COLUMN "public"."sys_permission"."component" IS '组件路径';
COMMENT ON COLUMN "public"."sys_permission"."extend" IS '扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单';
COMMENT ON COLUMN "public"."sys_permission"."remark" IS '备注';
COMMENT ON COLUMN "public"."sys_permission"."sort" IS '权重(排序)';
COMMENT ON COLUMN "public"."sys_permission"."status" IS '0=禁用 1=开启';
COMMENT ON COLUMN "public"."sys_permission"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."sys_permission"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."sys_permission"."deleted_at" IS '删除时间';

DROP TABLE IF EXISTS "public"."sys_role";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_role" (
    "id" uuid NOT NULL,
    "tenant_id" uuid,
    "pid" uuid NOT NULL,
    "name" varchar(50) NOT NULL,
    "permission_ids" text,
    "remark" varchar(200),
    "status" int2 NOT NULL,
    "sort" int8 NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_role"."id" IS '编号';
COMMENT ON COLUMN "public"."sys_role"."tenant_id" IS '租户ID';
COMMENT ON COLUMN "public"."sys_role"."pid" IS '父级id';
COMMENT ON COLUMN "public"."sys_role"."name" IS '名称';
COMMENT ON COLUMN "public"."sys_role"."permission_ids" IS '菜单权限集合';
COMMENT ON COLUMN "public"."sys_role"."remark" IS '备注';
COMMENT ON COLUMN "public"."sys_role"."status" IS '0=禁用 1=开启';
COMMENT ON COLUMN "public"."sys_role"."sort" IS '排序值';
COMMENT ON COLUMN "public"."sys_role"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."sys_role"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."sys_role"."deleted_at" IS '删除时间';

DROP TABLE IF EXISTS "public"."sys_tenant";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."sys_tenant" (
    "id" uuid,
    "name" varchar,
    "email" varchar,
    "phone" varchar,
    "address" varchar,
    "type" varchar,
    "status" int2,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz
);

-- Column Comment
COMMENT ON COLUMN "public"."sys_tenant"."name" IS '租户名称';
COMMENT ON COLUMN "public"."sys_tenant"."email" IS '租户联系邮箱';
COMMENT ON COLUMN "public"."sys_tenant"."phone" IS '租户联系电话';
COMMENT ON COLUMN "public"."sys_tenant"."address" IS '租户地址';
COMMENT ON COLUMN "public"."sys_tenant"."type" IS '租户类型';
COMMENT ON COLUMN "public"."sys_tenant"."status" IS '状态';



-- Table Comment
COMMENT ON TABLE "public"."sys_admin" IS '系统-用户';


-- Table Comment
COMMENT ON TABLE "public"."sys_api" IS '系统-接口';


-- Table Comment
COMMENT ON TABLE "public"."sys_dept" IS '系统-部门';


-- Table Comment
COMMENT ON TABLE "public"."sys_dict" IS '系统-参数';


-- Table Comment
COMMENT ON TABLE "public"."sys_job" IS '系统-工作岗位';


-- Table Comment
COMMENT ON TABLE "public"."sys_operation_log" IS '系统-日志';


-- Table Comment
COMMENT ON TABLE "public"."sys_permission" IS '菜单和权限规则表';


-- Table Comment
COMMENT ON TABLE "public"."sys_role" IS '系统-角色';


-- Table Comment
COMMENT ON TABLE "public"."sys_tenant" IS '系统-租户';
