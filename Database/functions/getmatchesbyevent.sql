CREATE OR REPLACE FUNCTION public.getmatchesbyevent(event integer)
 RETURNS TABLE(id integer, game text)
 LANGUAGE sql
AS $function$

	select m.id     as id
	     , g."name" as game
	from public.matches as m
	join public.games as g
	  on m.game = g.id
	where m.event = $1;

$function$;
