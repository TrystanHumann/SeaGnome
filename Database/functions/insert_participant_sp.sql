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


