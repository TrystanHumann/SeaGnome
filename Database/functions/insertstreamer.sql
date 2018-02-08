CREATE OR REPLACE FUNCTION public.insertstreamer(tag text, active boolean)
 RETURNS text
 LANGUAGE sql
AS $function$

insert into public.streamers(tag, active)
values($1, $2)
returning tag


 $function$
