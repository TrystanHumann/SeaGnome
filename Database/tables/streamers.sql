CREATE TABLE public.streamers (
	id serial NOT NULL,
	tag text NOT NULL,
	active bool NOT NULL DEFAULT true,
	CONSTRAINT streamers_pk PRIMARY KEY (id),
	CONSTRAINT streamers_unique_streamer UNIQUE (tag)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX streamers_id_idx ON streamers USING btree (id) ;
