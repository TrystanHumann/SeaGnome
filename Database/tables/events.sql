CREATE TABLE public.events (
	id smallserial NOT NULL,
	name text NOT NULL,
	complete bool NOT NULL,
	created_date timestamptz not NULL,
	active bool not null,
	CONSTRAINT events_pkey PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX events_pk_index ON events USING btree (id) ;
