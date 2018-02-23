-- DROP TABLE public.events CASCADE

CREATE TABLE "public".events (
	id smallserial NOT NULL,
	"name" text NOT NULL,
	complete boolean NOT NULL,
	created_date timestamp with time zone,
	PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX events_pk_index ON events USING btree (id) ;
