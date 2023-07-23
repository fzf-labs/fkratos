--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Debian 14.5-2.pgdg110+2)
-- Dumped by pg_dump version 15.3 (Homebrew)

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
-- Name: sys_tenant; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_tenant (
    id uuid,
    name character varying,
    email character varying,
    phone character varying,
    address character varying,
    type character varying,
    status smallint DEFAULT 0,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.sys_tenant OWNER TO postgres;

--
-- Name: TABLE sys_tenant; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sys_tenant IS '系统-租户';


--
-- Name: COLUMN sys_tenant.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.id IS 'ID';


--
-- Name: COLUMN sys_tenant.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.name IS '租户名称';


--
-- Name: COLUMN sys_tenant.email; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.email IS '租户联系邮箱';


--
-- Name: COLUMN sys_tenant.phone; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.phone IS '租户联系电话';


--
-- Name: COLUMN sys_tenant.address; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.address IS '租户地址';


--
-- Name: COLUMN sys_tenant.type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.type IS '租户类型';


--
-- Name: COLUMN sys_tenant.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.status IS '状态';


--
-- Name: COLUMN sys_tenant.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.created_at IS '创建时间';


--
-- Name: COLUMN sys_tenant.updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.updated_at IS '更新时间';


--
-- Name: COLUMN sys_tenant.deleted_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tenant.deleted_at IS '删除时间';


--
-- Name: sys_tenant_pkey; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX sys_tenant_pkey ON public.sys_tenant USING btree (id);


--
-- PostgreSQL database dump complete
--

