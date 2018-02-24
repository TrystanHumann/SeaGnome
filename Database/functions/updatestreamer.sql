CREATE OR REPLACE FUNCTION public.updatestreamer(id integer, active boolean)
 RETURNS text
 LANGUAGE sql
AS $function$

update public.streamers
set active = $2
where id = $1
returning tag

 $function$;
