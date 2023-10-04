--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Debian 14.5-2.pgdg110+2)
-- Dumped by pg_dump version 15.4 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: user_role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_role (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    pid uuid NOT NULL,
    type character varying(10) DEFAULT 'menu'::character varying NOT NULL,
    title character varying(50) NOT NULL,
    name character varying(50) NOT NULL,
    path character varying(100) NOT NULL,
    icon character varying(50),
    menu_type character varying(10) DEFAULT 'tab'::character varying,
    url character varying(255),
    component character varying(100),
    no_login_valid boolean,
    extend character varying(20) DEFAULT 'none'::character varying,
    remark character varying(255),
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.user_role OWNER TO postgres;

--
-- Name: TABLE user_role; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.user_role IS '用户规则表';


--
-- Name: COLUMN user_role.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.id IS 'ID';


--
-- Name: COLUMN user_role.pid; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.pid IS '上级菜单';


--
-- Name: COLUMN user_role.type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.type IS '类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮';


--
-- Name: COLUMN user_role.title; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.title IS '标题';


--
-- Name: COLUMN user_role.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.name IS '规则名称';


--
-- Name: COLUMN user_role.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.path IS '路由路径';


--
-- Name: COLUMN user_role.icon; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.icon IS '图标';


--
-- Name: COLUMN user_role.menu_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.menu_type IS '菜单类型:tab=选项卡,link=链接,iframe=Iframe';


--
-- Name: COLUMN user_role.url; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.url IS 'URL';


--
-- Name: COLUMN user_role.component; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.component IS '组件路径';


--
-- Name: COLUMN user_role.no_login_valid; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.no_login_valid IS '未登录有效:0=否,1=是';


--
-- Name: COLUMN user_role.extend; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.extend IS '扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单';


--
-- Name: COLUMN user_role.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.remark IS '备注';


--
-- Name: COLUMN user_role.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.status IS '状态';


--
-- Name: COLUMN user_role.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.created_at IS '创建时间';


--
-- Name: COLUMN user_role.updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.updated_at IS '更新时间';


--
-- Name: COLUMN user_role.deleted_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_role.deleted_at IS '删除时间';


--
-- Name: user_role user_role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_role
    ADD CONSTRAINT user_role_pkey PRIMARY KEY (id);


--
-- Name: user_role_pid_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_role_pid_idx ON public.user_role USING btree (pid);


--
-- PostgreSQL database dump complete
--

