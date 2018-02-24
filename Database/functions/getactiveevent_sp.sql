CREATE OR REPLACE FUNCTION public.getactiveevent_sp
()
 RETURNS TABLE
(id smallint, name text, complete boolean, created_date timestamp
with time zone, "activeflag" boolean)
 LANGUAGE plpgsql
AS $function$
begin
    return query
    select e.id as "id", e.name as "name", e.complete as "complete", e.created_date as "created_date", e.active as "activeflag"
    from public.events e
    where e.active = true;
end;
$function$
