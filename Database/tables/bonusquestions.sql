CREATE TABLE "public".bonusquestions (
	id serial NOT NULL,
	question text NOT NULL,
	event int4 NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (event) REFERENCES events(id) ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX bonusquestions_pk_index ON bonusquestions USING btree (id) ;
CREATE INDEX fki_bonusquestions_event_events_id_fk ON bonusquestions USING btree (event) ;
