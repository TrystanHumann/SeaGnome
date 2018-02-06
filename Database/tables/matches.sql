CREATE TABLE "public".matches (
	id serial NOT NULL,
	event int4 NOT NULL,
	game int4 NOT NULL,
	scheduled timestamptz NULL,
	completed timestamptz NULL,
	winner int4 NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (event) REFERENCES events(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	FOREIGN KEY (game) REFERENCES games(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	FOREIGN KEY (winner) REFERENCES competitors(id) ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
	OIDS=FALSE
) ;
CREATE INDEX fki_matches_event_events_id_fk ON matches USING btree (event) ;
CREATE INDEX fki_matches_game_games_id_fk ON matches USING btree (game) ;
CREATE INDEX fki_matches_winner_competitors_id_fk ON matches USING btree (winner) ;
CREATE UNIQUE INDEX matches_pk_index ON matches USING btree (id) ;
