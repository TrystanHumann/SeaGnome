CREATE TABLE public.bonusquestions (
	id serial NOT NULL,
	question text NOT NULL,
	event int4 NOT NULL,
	tiebreaker bool NOT NULL,
	CONSTRAINT bonusquestions_pkey PRIMARY KEY (id),
	CONSTRAINT bonusquestions_unique_question_event UNIQUE (question, event)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX bonusquestions_pk_index ON bonusquestions USING btree (id) ;
CREATE INDEX fki_bonusquestions_event_events_id_fk ON bonusquestions USING btree (event) ;
