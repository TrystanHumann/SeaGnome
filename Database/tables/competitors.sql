CREATE TABLE public.competitors (
	id smallserial NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT competitors_pk PRIMARY KEY (id),
	CONSTRAINT competitors_unique_competitors UNIQUE (name)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX competitors_pk_index ON competitors USING btree (id) ;
