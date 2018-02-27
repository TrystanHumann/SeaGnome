CREATE OR REPLACE FUNCTION public.getmatchesbyeventforresults(event integer)
 RETURNS TABLE(id integer, gameid integer, game text, competitorid smallint, competitor text)
 LANGUAGE sql
AS $function$

	select m.id     as id
	     , g.id     as gameid
	     , g."name" as gamename
	     , c.id     as competitorid
	     , c."name" as competitor
	from public.matches as m
	join public.games as g
	  on m.game = g.id
	join public.participants as p
	  on m.id = p."match"
	join public.competitors as c
	  on p.competitor = c.id
	where m.event = $1
	  and c."name" != 'Skip this'
	order by m.id;

$function$
