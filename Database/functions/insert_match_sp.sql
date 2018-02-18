-- drop function public.insert_match_sp(int, int, timestamp, /*timestamp,*/ int)
CREATE OR REPLACE FUNCTION public.insert_match_sp(ev int, ga int, sch timestamp = null, /*comp timestamp = null,*/ win int = null) returns table(returnid int)
 LANGUAGE plpgsql
AS $$
declare fetchID int;
begin
	-- Assuming they don't update competitors
	fetchID := (select m.id from public.matches m where m.event = ev and m.game = ga);
	if fetchID is null then
	-- insert
		insert into public.matches(event, game, scheduled, winner)
		values (ev, ga, sch, /*comp,*/ win);
	else 
		-- update 
		update public.matches
		set
			event = ev,
			game = ga, 
			scheduled = sch,
			--completed = comp, 
			winner = win  
		where m.event = event and m.game = game;	
	end if;
	return query
	select m.id as returnid from public.matches m where m.event = ev and m.game = ga;	
	--RAISE INFO 'Out variable: %', outUserID;
end $$;


