CREATE TABLE "public".games (
	id serial NOT NULL,
	"name" text NOT NULL,
	PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX games_pk_index ON games USING btree (id) ;
