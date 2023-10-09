CREATE TABLE public."user" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    uid character varying(64) NOT NULL,
    user_group_id character varying(64) DEFAULT ''::character varying NOT NULL,
    username character varying(64) DEFAULT ''::character varying NOT NULL,
    phone character varying DEFAULT ''::character varying NOT NULL,
    email character varying DEFAULT ''::character varying NOT NULL,
    password character varying DEFAULT ''::character varying NOT NULL,
    salt character varying DEFAULT ''::character varying NOT NULL,
    nickname character varying DEFAULT ''::character varying NOT NULL,
    sex smallint DEFAULT 0 NOT NULL,
    avatar character varying DEFAULT ''::character varying NOT NULL,
    profile character varying DEFAULT ''::character varying NOT NULL,
    other jsonb,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public."user" IS '用户表';
COMMENT ON COLUMN public."user".id IS 'Id';
COMMENT ON COLUMN public."user".uid IS 'uid';
COMMENT ON COLUMN public."user".user_group_id IS '分组ID';
COMMENT ON COLUMN public."user".username IS '用户名';
COMMENT ON COLUMN public."user".phone IS '手机';
COMMENT ON COLUMN public."user".email IS '邮箱';
COMMENT ON COLUMN public."user".password IS '密码';
COMMENT ON COLUMN public."user".salt IS '盐值';
COMMENT ON COLUMN public."user".nickname IS '昵称';
COMMENT ON COLUMN public."user".sex IS '性别（0未知 1男 2女）';
COMMENT ON COLUMN public."user".avatar IS '头像';
COMMENT ON COLUMN public."user".profile IS '简介';
COMMENT ON COLUMN public."user".other IS '其他';
COMMENT ON COLUMN public."user".status IS '状态';
COMMENT ON COLUMN public."user".created_at IS '创建时间';
COMMENT ON COLUMN public."user".updated_at IS '更新时间';
COMMENT ON COLUMN public."user".deleted_at IS '删除时间';
ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);
CREATE INDEX user_email_idx ON public."user" USING btree (email);
CREATE INDEX user_phone_idx ON public."user" USING btree (phone);
CREATE UNIQUE INDEX user_uid_idx ON public."user" USING btree (uid);
CREATE INDEX user_user_group_id_idx ON public."user" USING btree (user_group_id);
CREATE UNIQUE INDEX user_username_idx ON public."user" USING btree (username);
