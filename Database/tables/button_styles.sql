drop table if exists public.button_styles
CREATE TABLE "public".button_styles (
	button_id uuid NOT NULL,
	button_color varchar(18) NOT NULL,
	button_text varchar(50) NULL,
	button_link text NULL,
	CONSTRAINT button_styles_pk PRIMARY KEY (button_id)
)
WITH (
	OIDS=FALSE
);