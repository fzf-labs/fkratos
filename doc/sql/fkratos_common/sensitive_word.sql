CREATE TABLE public.sensitive_word (
    id uuid NOT NULL,
    word character varying NOT NULL,
    labs jsonb,
    "desc" character varying,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON COLUMN public.sensitive_word.id IS 'ID';
COMMENT ON COLUMN public.sensitive_word.word IS '敏感词';
COMMENT ON COLUMN public.sensitive_word.labs IS '标签';
COMMENT ON COLUMN public.sensitive_word."desc" IS '描述';
COMMENT ON COLUMN public.sensitive_word.created_at IS '创建时间';
COMMENT ON COLUMN public.sensitive_word.updated_at IS '更新时间';
COMMENT ON COLUMN public.sensitive_word.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sensitive_word
    ADD CONSTRAINT sensitive_word_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX sensitive_word_word_idx ON public.sensitive_word USING btree (word);
