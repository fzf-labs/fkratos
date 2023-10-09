CREATE TABLE public.user_group (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(50),
    roles jsonb,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_group IS '用户分组表';
COMMENT ON COLUMN public.user_group.id IS 'ID';
COMMENT ON COLUMN public.user_group.name IS '名称';
COMMENT ON COLUMN public.user_group.roles IS '权限';
COMMENT ON COLUMN public.user_group.status IS '状态';
COMMENT ON COLUMN public.user_group.created_at IS '创建时间';
COMMENT ON COLUMN public.user_group.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_group.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_group
    ADD CONSTRAINT user_group_pkey PRIMARY KEY (id);
