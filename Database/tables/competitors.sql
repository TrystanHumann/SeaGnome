CREATE TABLE "public".competitors (
	id smallserial NOT NULL,
	"name" text NOT NULL,
	PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX competitors_pk_index ON competitors USING btree (id) ;
