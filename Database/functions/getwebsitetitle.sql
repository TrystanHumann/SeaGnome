CREATE OR REPLACE FUNCTION public.getwebsitetitle()
 RETURNS TABLE(title varchar)
 LANGUAGE sql
AS $function$
    -- getting first title
	select title 
    from public.webpagetitle
    limit 1;
$function$;
