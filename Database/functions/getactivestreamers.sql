CREATE OR REPLACE FUNCTION public.getactivestreamers(max integer)
 RETURNS TABLE(id integer, tag text, active boolean)
 LANGUAGE sql
AS $function$

select id, tag, active
from public.streamers
where active = true
limit $1

 $function$;
