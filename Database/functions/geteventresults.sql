CREATE OR REPLACE FUNCTION public.geteventresults(event integer)
 RETURNS TABLE(name text, wins bigint, matches bigint)
 LANGUAGE sql
AS $function$

	select com."name"
	     , count(par.id) as wins
	     , (select count(p2.id) 
	          from public.participants as p2
	    	  join public.matches      as m2
			    on p2."match" = m2.id
			   and m2.winner is not null
			   and ((m2.event = $1) or ($1 = -1))
			  join public.competitors  as c2
			    on p2.competitor = c2.id
			   and c2."name" = com."name") as matches
	from public.participants as par
	join public.matches      as mat
	  on par."match" = mat.id
	 and mat.winner = par.competitor
	 and ((mat.event = $1) or ($1 = -1))
	join public.competitors  as com
	  on par.competitor = com.id
	group by com."name"

$function$;
