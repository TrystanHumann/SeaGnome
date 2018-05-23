CREATE OR REPLACE FUNCTION public.getevents_sp()
 RETURNS TABLE(id smallint, name text, complete boolean, created_date timestamp with time zone)
 LANGUAGE plpgsql
AS $function$
	begin
        return query 
        select e.id as "id", e.name as "name", e.complete as "complete", e.created_date as "created_date"
        from public.events e 
        order by e.created_date desc;
	end;
$function$;
