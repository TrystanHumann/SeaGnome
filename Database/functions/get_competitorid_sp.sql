CREATE OR REPLACE FUNCTION public.get_competitorid_sp
(competitor text)
 RETURNS TABLE
(returnid smallint)
 LANGUAGE plpgsql
AS $function$
begin

    return query
    select c.id as returnid
    from public.competitors c
    where c."name" = competitor;

end
$function$
