CREATE OR REPLACE FUNCTION public.getcompetitorsbymatch(matchsearch integer)
 RETURNS TABLE(id smallint, compname text)
 LANGUAGE sql
AS $function$

	select c.id     as id
	     , c."name" as compname
	from public.participants as p
	join public.competitors as c
	  on p.competitor = c.id
	where p."match" = $1;

$function$
