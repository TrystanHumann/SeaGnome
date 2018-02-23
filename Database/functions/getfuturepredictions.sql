CREATE OR REPLACE FUNCTION public.getfuturepredictions()
 RETURNS TABLE(game text, participant text, votes bigint)
 LANGUAGE sql
AS $function$
	select gam."name" as game, com."name" as participant, count(com.id) as votes
	from public.predictions as pre
	join public.users as use
	  on pre."user" = use.id
	join public.participants as par
	  on pre.participant = par.id
	join public.competitors as com
	  on par.competitor = com.id
	join public.matches as mat
	  on par."match" = mat.id
	join public.games as gam
	  on mat.game = gam.id
	--where mat.scheduled > now()
	group by com.id, gam."name", mat.scheduled
	order by mat.scheduled asc, count(com.id) desc
$function$
