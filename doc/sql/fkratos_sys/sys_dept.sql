CREATE TABLE public.sys_dept (
    id uuid NOT NULL,
    pid uuid NOT NULL,
    name character varying(50) NOT NULL,
    full_name character varying(50) NOT NULL,
    responsible character varying(20),
    phone character varying(20),
    email character varying(255),
    type smallint NOT NULL,
    status smallint NOT NULL,
    sort bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_dept IS '系统-部门';
COMMENT ON COLUMN public.sys_dept.id IS '编号';
COMMENT ON COLUMN public.sys_dept.pid IS '父级id';
COMMENT ON COLUMN public.sys_dept.name IS '部门简称';
COMMENT ON COLUMN public.sys_dept.full_name IS '部门全称';
COMMENT ON COLUMN public.sys_dept.responsible IS '负责人';
COMMENT ON COLUMN public.sys_dept.phone IS '负责人电话';
COMMENT ON COLUMN public.sys_dept.email IS '负责人邮箱';
COMMENT ON COLUMN public.sys_dept.type IS '1=公司 2=子公司 3=部门';
COMMENT ON COLUMN public.sys_dept.status IS '0=禁用 1=开启';
COMMENT ON COLUMN public.sys_dept.sort IS '排序值';
COMMENT ON COLUMN public.sys_dept.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_dept.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_dept.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_dept
    ADD CONSTRAINT sys_dept_pkey PRIMARY KEY (id);
