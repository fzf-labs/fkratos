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
-- Name: dict_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dict_data (
    id uuid NOT NULL,
    type character varying(100) NOT NULL,
    label character varying(100) NOT NULL,
    value character varying(100) NOT NULL,
    remark character varying(500),
    css_color character varying(100),
    css_class character varying(100),
    status smallint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.dict_data OWNER TO postgres;

--
-- Name: TABLE dict_data; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.dict_data IS '字典数据表';


--
-- Name: COLUMN dict_data.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.id IS 'id';


--
-- Name: COLUMN dict_data.type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.type IS '字典类型';


--
-- Name: COLUMN dict_data.label; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.label IS '字典标签';


--
-- Name: COLUMN dict_data.value; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.value IS '字典键值';


--
-- Name: COLUMN dict_data.remark; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.remark IS '备注';


--
-- Name: COLUMN dict_data.css_color; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.css_color IS 'css 颜色';


--
-- Name: COLUMN dict_data.css_class; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.css_class IS 'css 样式';


--
-- Name: COLUMN dict_data.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.status IS '状态（0正常 1停用）';


--
-- Name: COLUMN dict_data.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.created_at IS '创建时间';


--
-- Name: COLUMN dict_data.updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.updated_at IS '更新时间';


--
-- Name: COLUMN dict_data.deleted_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.dict_data.deleted_at IS '删除时间';


--
-- Name: dict_data dict_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dict_data
    ADD CONSTRAINT dict_data_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

