CREATE TABLE "public".predictions (
	id serial NOT NULL,
	"user" int4 NOT NULL,
	participant int4 NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (participant) REFERENCES participants(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	FOREIGN KEY ("user") REFERENCES users(id) ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
	OIDS=FALSE
) ;
CREATE INDEX fki_predictions_participant_participants_id_fk ON predictions USING btree (participant) ;
CREATE INDEX fki_predictions_user_users_id_fk ON predictions USING btree ("user") ;
CREATE UNIQUE INDEX predictions_id_pk ON predictions USING btree (id) ;
