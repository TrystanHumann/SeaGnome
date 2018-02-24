CREATE TABLE public.auth (
	id bytea NOT NULL,
	username text NOT NULL,
	password text NOT NULL,
	salt text NOT NULL,
	token bytea NULL,
	expires timestamptz NULL,
	CONSTRAINT auth_pk PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;
