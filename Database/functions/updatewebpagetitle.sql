CREATE OR REPLACE FUNCTION public.updateTitle(varchar)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
declare 
	newtitle varchar(100) = $1;
BEGIN
    -- setting title for them all
    update public.webpagetitle
    set title = newtitle
END
$function$;
