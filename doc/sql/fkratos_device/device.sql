CREATE TABLE public.device (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    "SNId" character varying(255) NOT NULL,
    "deviceBrand" character varying(64),
    "modelDevice" character varying(125),
    status smallint,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.device IS '设备表';
COMMENT ON COLUMN public.device.id IS '记录ID';
COMMENT ON COLUMN public.device."SNId" IS '设备SN唯一标识码';
COMMENT ON COLUMN public.device."deviceBrand" IS '设备品牌';
COMMENT ON COLUMN public.device."modelDevice" IS '设备型号';
COMMENT ON COLUMN public.device.status IS '状态';
COMMENT ON COLUMN public.device.created_at IS '创建时间';
COMMENT ON COLUMN public.device.updated_at IS '更新时间';
COMMENT ON COLUMN public.device.deleted_at IS '删除时间';
ALTER TABLE ONLY public.device
    ADD CONSTRAINT devices_pkey PRIMARY KEY (id);
