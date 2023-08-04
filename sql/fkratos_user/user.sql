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
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id uuid NOT NULL,
    username character varying,
    phone character varying,
    email character varying,
    password character varying NOT NULL,
    salt character varying NOT NULL,
    nickname character varying DEFAULT ''::character varying,
    sex smallint DEFAULT 0,
    avatar character varying DEFAULT ''::character varying,
    profile character varying,
    other json,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: COLUMN "user".id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".id IS 'Id';


--
-- Name: COLUMN "user".username; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".username IS '用户名';


--
-- Name: COLUMN "user".phone; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".phone IS '手机';


--
-- Name: COLUMN "user".email; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".email IS '邮箱';


--
-- Name: COLUMN "user".password; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".password IS '密码';


--
-- Name: COLUMN "user".salt; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".salt IS '盐值';


--
-- Name: COLUMN "user".nickname; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".nickname IS '昵称';


--
-- Name: COLUMN "user".sex; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".sex IS '性别（0未知 1男 2女）';


--
-- Name: COLUMN "user".avatar; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".avatar IS '头像';


--
-- Name: COLUMN "user".profile; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".profile IS '简介';


--
-- Name: COLUMN "user".other; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".other IS '其他';


--
-- Name: COLUMN "user".status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".status IS '状态';


--
-- Name: COLUMN "user".created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".created_at IS '创建时间';


--
-- Name: COLUMN "user".updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".updated_at IS '更新时间';


--
-- Name: COLUMN "user".deleted_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".deleted_at IS '删除时间';


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: user_email_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_email_idx ON public."user" USING btree (email);


--
-- Name: user_phone_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_phone_idx ON public."user" USING btree (phone);


--
-- Name: user_status_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_status_idx ON public."user" USING btree (status);


--
-- Name: user_username_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_username_idx ON public."user" USING btree (username);


--
-- PostgreSQL database dump complete
--

