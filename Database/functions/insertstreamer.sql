CREATE OR REPLACE FUNCTION public.insertstreamer
(streamerone text, streamertwo text)
 RETURNS void
 LANGUAGE sql
AS $function$

truncate public.streamers;

insert into public.streamers
	(tag, active)
values($1, true);

insert into public.streamers
	(tag, active)
values($2, true);


$function$
