CREATE OR REPLACE FUNCTION public.insert_match_sp(ev integer, ga integer, sch timestamp without time zone DEFAULT NULL::timestamp without time zone, win integer DEFAULT NULL::integer)
 RETURNS TABLE(returnid integer)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
	-- insert
	insert into public.matches(event, game, scheduled, winner)
	values (ev, ga, sch, win)
	on conflict(event, game) do update 
	set
		event = ev,
		game = ga, 
		scheduled = coalesce(sch, (select scheduled from public.matches where event = ev and game = ga)),
		winner = coalesce(nullif(win, 0), (select winner from public.matches where event = ev and game = ga));
	return query
	select m.id as returnid from public.matches m where m.event = ev and m.game = ga;	
	--RAISE INFO 'Out variable: %', outUserID;
end $function$;
