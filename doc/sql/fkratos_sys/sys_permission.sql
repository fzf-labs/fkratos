CREATE TABLE public.sys_permission (
    id uuid NOT NULL,
    pid uuid NOT NULL,
    type character varying(255) NOT NULL,
    title character varying(50) NOT NULL,
    name character varying(50) NOT NULL,
    path character varying(100) NOT NULL,
    icon character varying(50) NOT NULL,
    menu_type character varying(255),
    url character varying(255) NOT NULL,
    component character varying(100) NOT NULL,
    extend character varying(255) NOT NULL,
    remark character varying(255) NOT NULL,
    sort bigint NOT NULL,
    status smallint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_permission IS '菜单和权限规则表';
COMMENT ON COLUMN public.sys_permission.pid IS '上级菜单';
COMMENT ON COLUMN public.sys_permission.type IS '类型:menu_dir=菜单目录,menu=菜单项,button=页面按钮';
COMMENT ON COLUMN public.sys_permission.title IS '标题';
COMMENT ON COLUMN public.sys_permission.name IS '规则名称';
COMMENT ON COLUMN public.sys_permission.path IS '路由路径';
COMMENT ON COLUMN public.sys_permission.icon IS '图标';
COMMENT ON COLUMN public.sys_permission.menu_type IS '菜单类型:tab=选项卡,link=链接,iframe=Iframe';
COMMENT ON COLUMN public.sys_permission.url IS 'Url';
COMMENT ON COLUMN public.sys_permission.component IS '组件路径';
COMMENT ON COLUMN public.sys_permission.extend IS '扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单';
COMMENT ON COLUMN public.sys_permission.remark IS '备注';
COMMENT ON COLUMN public.sys_permission.sort IS '权重(排序)';
COMMENT ON COLUMN public.sys_permission.status IS '0=禁用 1=开启';
COMMENT ON COLUMN public.sys_permission.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_permission.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_permission.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_permission
    ADD CONSTRAINT sys_permission_pkey PRIMARY KEY (id);
