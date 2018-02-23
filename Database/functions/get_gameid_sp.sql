CREATE OR REPLACE FUNCTION public.get_gameid_sp
(gamename text)
 RETURNS TABLE
(returnid integer)
 LANGUAGE plpgsql
AS $function$
begin
    return query
    select g.id as returnid
    from public.games g
    where g."name" = gamename;
end
$function$
