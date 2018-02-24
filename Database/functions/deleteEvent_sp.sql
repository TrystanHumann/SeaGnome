CREATE OR REPLACE FUNCTION public.deleteevent_sp(smallint)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
declare idParam int2 = $1;
BEGIN
    delete from public.events
    where id = idParam;
END
$function$
