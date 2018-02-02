CREATE TABLE "public".events (
	id smallserial NOT NULL,
	"name" text NOT NULL,
	complete bit(1) NOT NULL,
	PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX events_pk_index ON events USING btree (id) ;
