CREATE TABLE public.events (
	id smallserial NOT NULL,
	"name" text NOT NULL,
	complete bool NOT NULL,
	created_date timestamptz NOT NULL,
	active bool NOT NULL,
	CONSTRAINT events_pkey PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX events_pk_index ON events USING btree (id) ;
