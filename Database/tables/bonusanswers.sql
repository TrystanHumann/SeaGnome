CREATE TABLE public.bonusanswers (
	id serial NOT NULL,
	question int4 NOT NULL,
	answer text NOT NULL,
	CONSTRAINT bonusanswers_pkey PRIMARY KEY (id),
	CONSTRAINT bonusanswers_question_fkey FOREIGN KEY (question) REFERENCES bonusquestions(id) ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX bonusanswers_pk_index ON bonusanswers USING btree (id) ;
CREATE INDEX fki_bonusanswers_question_to_bonusquestions_fk ON bonusanswers USING btree (question) ;
