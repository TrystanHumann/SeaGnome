CREATE TABLE public.users (
	id serial NOT NULL,
	twitchun text NOT NULL,
	twitterhandle text NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique_twitchun UNIQUE (twitchun)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX users_pk_index ON users USING btree (id) ;
