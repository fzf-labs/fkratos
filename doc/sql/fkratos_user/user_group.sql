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
-- Name: user_group; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_group (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(50),
    roles jsonb,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.user_group OWNER TO postgres;

--
-- Name: TABLE user_group; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.user_group IS '用户分组表';


--
-- Name: COLUMN user_group.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_group.id IS 'ID';


--
-- Name: COLUMN user_group.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_group.name IS '名称';


--
-- Name: COLUMN user_group.roles; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_group.roles IS '权限';


--
-- Name: COLUMN user_group.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_group.status IS '状态';


--
-- Name: COLUMN user_group.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_group.created_at IS '创建时间';


--
-- Name: COLUMN user_group.updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_group.updated_at IS '更新时间';


--
-- Name: COLUMN user_group.deleted_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.user_group.deleted_at IS '删除时间';


--
-- Name: user_group user_group_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_group
    ADD CONSTRAINT user_group_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

