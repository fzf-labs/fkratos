CREATE TABLE public.sys_role (
    id uuid NOT NULL,
    pid uuid NOT NULL,
    name character varying(50) NOT NULL,
    permission_ids text,
    remark character varying(200),
    status smallint NOT NULL,
    sort bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_role IS '系统-角色';
COMMENT ON COLUMN public.sys_role.id IS '编号';
COMMENT ON COLUMN public.sys_role.pid IS '父级id';
COMMENT ON COLUMN public.sys_role.name IS '名称';
COMMENT ON COLUMN public.sys_role.permission_ids IS '菜单权限集合';
COMMENT ON COLUMN public.sys_role.remark IS '备注';
COMMENT ON COLUMN public.sys_role.status IS '0=禁用 1=开启';
COMMENT ON COLUMN public.sys_role.sort IS '排序值';
COMMENT ON COLUMN public.sys_role.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_role.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_role.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_role
    ADD CONSTRAINT sys_role_pkey PRIMARY KEY (id);
