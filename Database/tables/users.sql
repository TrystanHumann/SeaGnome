-- Table: public.users

-- DROP TABLE public.users;

CREATE TABLE public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    twitchun text COLLATE pg_catalog."default" NOT NULL,
    twitterhandle text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to development;

-- Index: users_pk_index

-- DROP INDEX public.users_pk_index;

CREATE UNIQUE INDEX users_pk_index
    ON public.users USING btree
    (id)
    TABLESPACE pg_default;

ALTER TABLE public.users
    CLUSTER ON users_pk_index;