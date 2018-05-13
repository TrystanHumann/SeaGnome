-- drop table if exists public.button_colors;
CREATE TABLE "public".button_colors (
	button_guid uuid NOT NULL,
	hex_color varchar(7) NOT NULL,
	PRIMARY KEY (button_guid),
	UNIQUE (button_guid)
)
WITH (
	OIDS=FALSE
) ;
