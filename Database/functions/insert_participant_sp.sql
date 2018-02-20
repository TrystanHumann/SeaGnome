<<<<<<< HEAD
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
=======
-- drop function public.insert_participant_sp(int,int)
CREATE OR REPLACE FUNCTION public.insert_participant_sp(matchid int, competitorid int) returns table(returnid int)
 LANGUAGE plpgsql
AS $$
declare fetchID int;
begin
	-- Assuming they don't update competitors
	fetchID := (select p.id from public.participants p where p."match" = matchid and p.competitor = competitorid);
	if fetchID is null then
	-- insert
		insert into public.participants(match, competitor)
		values (matchid, competitorid);
	end if;
	return query
	select p.id as returnid from public.participants p where p."match" = matchid and p.competitor = competitorid;
	--RAISE INFO 'Out variable: %', outUserID;
end $$;


>>>>>>> d305009f3ebb93b3f7e7c23c87f02ba6be5d6fd7
