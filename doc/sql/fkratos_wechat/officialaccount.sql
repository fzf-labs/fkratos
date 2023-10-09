CREATE TABLE public.officialaccount (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    app_id character varying(100) DEFAULT ''::character varying NOT NULL,
    app_secret character varying(100) DEFAULT ''::character varying NOT NULL,
    account_id character varying(11) DEFAULT ''::character varying NOT NULL,
    business_id character varying(11) DEFAULT ''::character varying NOT NULL,
    name character varying(50) DEFAULT ''::character varying NOT NULL,
    expires_access_token character varying(11) DEFAULT ''::character varying NOT NULL,
    expires_jsapi_ticket character varying(11) DEFAULT ''::character varying NOT NULL,
    access_token character varying(20) DEFAULT ''::character varying NOT NULL,
    jsapi_ticket character varying(255) DEFAULT ''::character varying NOT NULL,
    qrcode character varying(255) DEFAULT ''::character varying NOT NULL,
    token character varying(255) DEFAULT ''::character varying NOT NULL,
    encoding_ase_key character varying(255) DEFAULT ''::character varying NOT NULL,
    remark character varying(100) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.officialaccount IS '微信公众号配置表';
COMMENT ON COLUMN public.officialaccount.id IS 'ID';
COMMENT ON COLUMN public.officialaccount.app_id IS '开发者ID(AppID)';
COMMENT ON COLUMN public.officialaccount.app_secret IS '开发者密码(AppSecret)';
COMMENT ON COLUMN public.officialaccount.account_id IS '账号id';
COMMENT ON COLUMN public.officialaccount.business_id IS '业务主账号id';
COMMENT ON COLUMN public.officialaccount.name IS '公众号名称';
COMMENT ON COLUMN public.officialaccount.expires_access_token IS '获取access_token时间';
COMMENT ON COLUMN public.officialaccount.expires_jsapi_ticket IS '获取jsapi_ticket时间';
COMMENT ON COLUMN public.officialaccount.access_token IS 'access_token';
COMMENT ON COLUMN public.officialaccount.jsapi_ticket IS 'jsapi_ticket';
COMMENT ON COLUMN public.officialaccount.qrcode IS '二维码';
COMMENT ON COLUMN public.officialaccount.token IS 'Token 长度为3-32字符';
COMMENT ON COLUMN public.officialaccount.encoding_ase_key IS '消息加密密钥由43位字符组成';
COMMENT ON COLUMN public.officialaccount.remark IS '备注';
COMMENT ON COLUMN public.officialaccount.created_at IS '创建时间';
COMMENT ON COLUMN public.officialaccount.updated_at IS '更新时间';
COMMENT ON COLUMN public.officialaccount.deleted_at IS '删除时间';
ALTER TABLE ONLY public.officialaccount
    ADD CONSTRAINT officialaccount_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX officialaccount_app_id_idx ON public.officialaccount USING btree (app_id);
