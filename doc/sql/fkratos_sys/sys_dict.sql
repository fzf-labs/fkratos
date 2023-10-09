CREATE TABLE public.sys_dict (
    id uuid NOT NULL,
    pid uuid NOT NULL,
    name character varying(50) NOT NULL,
    type smallint NOT NULL,
    unique_key character varying(50) NOT NULL,
    value character varying(2048) NOT NULL,
    status smallint NOT NULL,
    sort numeric(20,0) NOT NULL,
    remark character varying(200) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_dict IS '系统-参数';
COMMENT ON COLUMN public.sys_dict.id IS '编号';
COMMENT ON COLUMN public.sys_dict.pid IS '0=配置集 !0=父级id';
COMMENT ON COLUMN public.sys_dict.name IS '名称';
COMMENT ON COLUMN public.sys_dict.type IS '1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单图 10多图 11单文件 12多文件';
COMMENT ON COLUMN public.sys_dict.unique_key IS '唯一值';
COMMENT ON COLUMN public.sys_dict.value IS '配置值';
COMMENT ON COLUMN public.sys_dict.status IS '0=禁用 1=开启';
COMMENT ON COLUMN public.sys_dict.sort IS '排序值';
COMMENT ON COLUMN public.sys_dict.remark IS '备注';
COMMENT ON COLUMN public.sys_dict.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_dict.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_dict.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_dict
    ADD CONSTRAINT sys_dict_pkey PRIMARY KEY (id);
