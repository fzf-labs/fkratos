CREATE TABLE public.miniprogram (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    app_id character varying(100) DEFAULT ''::character varying NOT NULL,
    app_secret character varying(100) DEFAULT ''::character varying NOT NULL,
    account_id character varying(11) DEFAULT ''::character varying NOT NULL,
    business_id character varying(11) DEFAULT ''::character varying NOT NULL,
    name character varying(50) DEFAULT ''::character varying NOT NULL,
    expires_access_token character varying(11) DEFAULT ''::character varying NOT NULL,
    expires_jsapi_ticket character varying(11) DEFAULT ''::character varying NOT NULL,
    qrcode character varying(255) DEFAULT ''::character varying NOT NULL,
    token character varying(255) DEFAULT ''::character varying NOT NULL,
    encoding_ase_key character varying(255) DEFAULT ''::character varying NOT NULL,
    remark character varying(100) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.miniprogram IS '微信小程序配置表';
COMMENT ON COLUMN public.miniprogram.id IS 'ID';
COMMENT ON COLUMN public.miniprogram.app_id IS '开发者ID(AppID)';
COMMENT ON COLUMN public.miniprogram.app_secret IS '开发者密码(AppSecret)';
COMMENT ON COLUMN public.miniprogram.account_id IS '账号id';
COMMENT ON COLUMN public.miniprogram.business_id IS '业务主账号id';
COMMENT ON COLUMN public.miniprogram.name IS '小程序名称';
COMMENT ON COLUMN public.miniprogram.expires_access_token IS '获取access_token时间';
COMMENT ON COLUMN public.miniprogram.expires_jsapi_ticket IS '获取jsapi_ticket时间';
COMMENT ON COLUMN public.miniprogram.qrcode IS '二维码';
COMMENT ON COLUMN public.miniprogram.token IS 'Token 长度为3-32字符';
COMMENT ON COLUMN public.miniprogram.encoding_ase_key IS '消息加密密钥由43位字符组成';
COMMENT ON COLUMN public.miniprogram.remark IS '备注';
COMMENT ON COLUMN public.miniprogram.created_at IS '创建时间';
COMMENT ON COLUMN public.miniprogram.updated_at IS '更新时间';
COMMENT ON COLUMN public.miniprogram.deleted_at IS '删除时间';
ALTER TABLE ONLY public.miniprogram
    ADD CONSTRAINT miniprogram_pkey PRIMARY KEY (id);
