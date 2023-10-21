CREATE TABLE public.device (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    sn character varying NOT NULL,
    device_name character varying NOT NULL,
    device_type character varying NOT NULL,
    device_model character varying NOT NULL,
    "desc" character varying,
    certificate character varying NOT NULL,
    secure_key character varying NOT NULL,
    firmware_version character varying,
    software_version character varying,
    registry_time timestamp with time zone,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.device IS '设备表';
COMMENT ON COLUMN public.device.id IS '记录ID';
COMMENT ON COLUMN public.device.sn IS '设备的唯一标识序列号';
COMMENT ON COLUMN public.device.device_name IS '设备名称';
COMMENT ON COLUMN public.device.device_type IS '设备类型';
COMMENT ON COLUMN public.device.device_model IS '设备型号';
COMMENT ON COLUMN public.device."desc" IS '描述';
COMMENT ON COLUMN public.device.certificate IS '设备证书';
COMMENT ON COLUMN public.device.secure_key IS '设备密钥';
COMMENT ON COLUMN public.device.firmware_version IS '固件版本号';
COMMENT ON COLUMN public.device.software_version IS '软件版本号';
COMMENT ON COLUMN public.device.registry_time IS '激活时间';
COMMENT ON COLUMN public.device.status IS '状态';
COMMENT ON COLUMN public.device.created_at IS '创建时间';
COMMENT ON COLUMN public.device.updated_at IS '更新时间';
COMMENT ON COLUMN public.device.deleted_at IS '删除时间';
ALTER TABLE ONLY public.device
    ADD CONSTRAINT device_pkey PRIMARY KEY (id);
CREATE INDEX device_sn_idx ON public.device USING btree (sn);
