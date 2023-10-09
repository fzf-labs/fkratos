CREATE TABLE public.sys_job (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    code character varying(32),
    remark character varying(255),
    sort bigint NOT NULL,
    status smallint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_job IS '系统-工作岗位';
COMMENT ON COLUMN public.sys_job.id IS '编号';
COMMENT ON COLUMN public.sys_job.name IS '岗位名称';
COMMENT ON COLUMN public.sys_job.code IS '岗位编码';
COMMENT ON COLUMN public.sys_job.remark IS '备注';
COMMENT ON COLUMN public.sys_job.sort IS '排序值';
COMMENT ON COLUMN public.sys_job.status IS '0=禁用 1=开启 ';
COMMENT ON COLUMN public.sys_job.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_job.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_job.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_job
    ADD CONSTRAINT sys_job_pkey PRIMARY KEY (id);
