CREATE OR REPLACE FUNCTION public.updateevent_sp(smallint, boolean)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
declare 
	idParam int2 = $1;
	compParam boolean = $2;
BEGIN
    update public.events
    set complete = compParam
    where id = idParam;
END
$function$;
