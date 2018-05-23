CREATE TABLE public.predictions (
	id serial NOT NULL,
	"user" int4 NOT NULL,
	participant int4 NOT NULL,
	CONSTRAINT predictions_pkey PRIMARY KEY (id),
	CONSTRAINT predictions_unique_user_participant UNIQUE ("user", participant),
	CONSTRAINT predictions_participant_participants_id_fk FOREIGN KEY (participant) REFERENCES participants(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT predictions_user_users_id_fk FOREIGN KEY ("user") REFERENCES users(id) ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
	OIDS=FALSE
) ;
CREATE INDEX fki_predictions_participant_participants_id_fk ON predictions USING btree (participant) ;
CREATE INDEX fki_predictions_user_users_id_fk ON predictions USING btree ("user") ;
CREATE UNIQUE INDEX predictions_id_pk ON predictions USING btree (id) ;