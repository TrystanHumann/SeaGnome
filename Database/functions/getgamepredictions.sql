CREATE OR REPLACE FUNCTION public.getgamepredictions(event integer, "user" integer)
 RETURNS TABLE(game text, prediction text, winner text)
 LANGUAGE sql
AS $function$
	select gam."name" as game
	     , com."name" as prediction
	     , coalesce((select "name" from public.competitors where id = mat.winner), '') as winner
	from public.predictions  as pre
	join public.participants as par
	  on pre.participant = par.id
	join public.competitors  as com
	  on par.competitor = com.id
	join public.matches      as mat
	  on par."match" = mat.id
	 and ((mat.event = $1) or ($1 = -1))
	join public.games        as gam
	  on mat.game = gam.id
	where ((pre."user" = $2) or ($2 = -1))
	order by mat.scheduled asc
$function$;
