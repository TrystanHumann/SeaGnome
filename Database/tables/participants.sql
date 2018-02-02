CREATE TABLE "public".participants (
	id serial NOT NULL,
	"match" int4 NOT NULL,
	competitor int4 NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (match) REFERENCES matches(id),
	FOREIGN KEY (competitor) REFERENCES competitors(id)
)
WITH (
	OIDS=FALSE
) ;
CREATE INDEX fki_particiants_match_match_id_fk ON participants USING btree (match) ;
CREATE INDEX fki_participant_competitor_competitors_id_fk ON participants USING btree (competitor) ;
CREATE UNIQUE INDEX participants_pk_index ON participants USING btree (id) ;
