-- drop table if exists public.button_colors;
CREATE TABLE "public".button_colors (
	button_guid uuid NOT NULL,
	hex_color varchar(7) NOT NULL
)
WITH (
	OIDS=FALSE
) ;
COMMENT ON TABLE "public".button_colors IS 'Table used to store button color' ;