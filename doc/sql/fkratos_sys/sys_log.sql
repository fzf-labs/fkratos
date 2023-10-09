CREATE TABLE public.sys_log (
    id uuid NOT NULL,
    admin_id uuid NOT NULL,
    ip character varying(32) NOT NULL,
    uri character varying(200) NOT NULL,
    useragent character varying(255),
    header json,
    req json,
    resp json,
    created_at timestamp with time zone NOT NULL
);
COMMENT ON TABLE public.sys_log IS '系统-日志';
COMMENT ON COLUMN public.sys_log.id IS '编号';
COMMENT ON COLUMN public.sys_log.admin_id IS '管理员ID';
COMMENT ON COLUMN public.sys_log.ip IS 'ip';
COMMENT ON COLUMN public.sys_log.uri IS '请求路径';
COMMENT ON COLUMN public.sys_log.useragent IS '浏览器标识';
COMMENT ON COLUMN public.sys_log.header IS 'header';
COMMENT ON COLUMN public.sys_log.req IS '请求数据';
COMMENT ON COLUMN public.sys_log.resp IS '响应数据';
COMMENT ON COLUMN public.sys_log.created_at IS '创建时间';
ALTER TABLE ONLY public.sys_log
    ADD CONSTRAINT sys_log_pkey PRIMARY KEY (id);
