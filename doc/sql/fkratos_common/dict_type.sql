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
-- Name: dict_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dict_type (
    id uuid NOT NULL,
    name character varying(100) NOT NULL,
    type character varying(100) NOT NULL,
    status smallint NOT NULL,
    remark character varying(500),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.dict_type OWNER TO postgres;

--
-- Name: TABLE dict_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.dict_type IS '字典类型表';


--
-- Name: COLUMN dict_type.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_type.id IS 'id';


--
-- Name: COLUMN dict_type.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_type.name IS '字典名称';


--
-- Name: COLUMN dict_type.type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_type.type IS '字典类型';


--
-- Name: COLUMN dict_type.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_type.status IS '状态（0正常 1停用）';


--
-- Name: COLUMN dict_type.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_type.remark IS '备注';


--
-- Name: COLUMN dict_type.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_type.created_at IS '创建时间';


--
-- Name: COLUMN dict_type.updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_type.updated_at IS '更新时间';


--
-- Name: COLUMN dict_type.deleted_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_type.deleted_at IS '删除时间';


--
-- Name: dict_type dict_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dict_type
    ADD CONSTRAINT dict_type_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

