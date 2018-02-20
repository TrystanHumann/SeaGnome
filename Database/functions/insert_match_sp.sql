<<<<<<< HEAD
CREATE OR REPLACE FUNCTION public.insert_match_sp
(ev integer, ga integer, sch timestamp without time zone DEFAULT NULL::timestamp without time zone, win integer DEFAULT NULL::integer)
 RETURNS TABLE
(returnid integer)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
    -- insert
    insert into public.matches
        (event, game, scheduled, winner)
    values
        (ev, ga, sch, win)
    on conflict
    (event, game) do
    update 
	set
		event = ev,
		game = ga, 
		scheduled = coalesce(sch, (select scheduled from public.matches where event = ev and game = ga)),
		winner = coalesce(win, (select winner from public.matches where event = ev and game = ga));
    return query
    select m.id as returnid
    from public.matches m
    where m.event = ev and m.game = ga;
--RAISE INFO 'Out variable: %', outUserID;
end
$function$
=======
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


>>>>>>> d305009f3ebb93b3f7e7c23c87f02ba6be5d6fd7
