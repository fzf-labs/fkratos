CREATE TABLE public.user_rule (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    pid character varying(50) DEFAULT ''::character varying NOT NULL,
    type character varying(10) DEFAULT 'menu'::character varying NOT NULL,
    title character varying(50) NOT NULL,
    name character varying(50) NOT NULL,
    path character varying(100) NOT NULL,
    icon character varying(50) DEFAULT ''::character varying NOT NULL,
    menu_type character varying(10) DEFAULT 'tab'::character varying NOT NULL,
    url character varying(255) DEFAULT ''::character varying NOT NULL,
    component character varying(100) DEFAULT ''::character varying NOT NULL,
    extend character varying(20) DEFAULT 'none'::character varying NOT NULL,
    remark character varying(255) DEFAULT ''::character varying NOT NULL,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_rule IS '用户规则表';
COMMENT ON COLUMN public.user_rule.id IS 'ID';
COMMENT ON COLUMN public.user_rule.pid IS '上级菜单';
COMMENT ON COLUMN public.user_rule.type IS '类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮';
COMMENT ON COLUMN public.user_rule.title IS '标题';
COMMENT ON COLUMN public.user_rule.name IS '规则名称';
COMMENT ON COLUMN public.user_rule.path IS '路由路径';
COMMENT ON COLUMN public.user_rule.icon IS '图标';
COMMENT ON COLUMN public.user_rule.menu_type IS '菜单类型:tab=选项卡,link=链接,iframe=Iframe';
COMMENT ON COLUMN public.user_rule.url IS 'URL';
COMMENT ON COLUMN public.user_rule.component IS '组件路径';
COMMENT ON COLUMN public.user_rule.extend IS '扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单';
COMMENT ON COLUMN public.user_rule.remark IS '备注';
COMMENT ON COLUMN public.user_rule.status IS '状态';
COMMENT ON COLUMN public.user_rule.created_at IS '创建时间';
COMMENT ON COLUMN public.user_rule.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_rule.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_rule
    ADD CONSTRAINT user_rule_pkey PRIMARY KEY (id);
CREATE INDEX user_rule_pid_idx ON public.user_rule USING btree (pid);
