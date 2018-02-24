CREATE TABLE public.games (
	id serial NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT games_pk PRIMARY KEY (id),
	CONSTRAINT games_unique_game_name UNIQUE (name)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX games_pk_index ON games USING btree (id) ;
