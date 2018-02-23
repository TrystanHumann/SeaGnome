CREATE OR REPLACE FUNCTION public.insert_participant_sp
(matchid integer, competitorid integer)
 RETURNS TABLE
(returnid integer)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
    -- insert
    insert into public.participants
        (match, competitor)
    values
        (matchid, competitorid)
    on conflict
    ("match", competitor) 
		do
    update
		set
		"matchid",
		competitor
    return query
    select p.id as returnid
    from public.participants p
    where p."match" = matchid and p.competitor = competitorid;
--RAISE INFO 'Out variable: %', outUserID;
end
$function$
