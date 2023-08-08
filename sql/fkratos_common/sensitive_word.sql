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
-- Name: sensitive_word; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sensitive_word (
    id uuid NOT NULL,
    word character varying,
    labs json,
    "desc" character varying,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.sensitive_word OWNER TO postgres;

--
-- Name: COLUMN sensitive_word.id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sensitive_word.id IS 'ID';


--
-- Name: COLUMN sensitive_word.word; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sensitive_word.word IS '敏感词';


--
-- Name: COLUMN sensitive_word.labs; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sensitive_word.labs IS '标签';


--
-- Name: COLUMN sensitive_word."desc"; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sensitive_word."desc" IS '描述';


--
-- Name: COLUMN sensitive_word.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sensitive_word.created_at IS '创建时间';


--
-- Name: COLUMN sensitive_word.updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sensitive_word.updated_at IS '更新时间';


--
-- Name: COLUMN sensitive_word.deleted_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sensitive_word.deleted_at IS '删除时间';


--
-- Name: sensitive_word sensitive_word_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sensitive_word
    ADD CONSTRAINT sensitive_word_pkey PRIMARY KEY (id);


--
-- Name: sensitive_word_word_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX sensitive_word_word_idx ON public.sensitive_word USING btree (word);


--
-- PostgreSQL database dump complete
--

