CREATE OR REPLACE FUNCTION public.getuserscore(event integer, usr integer)
 RETURNS TABLE("user" text, total bigint, percent numeric)
 LANGUAGE sql
AS $function$

	select users.twitchun
	     , count(preds.id) as correct
	     , round(100 * (cast (count(preds.id) as numeric)/(select count(matc.id)
			                                               from predictions as pred
														   join participants as part
															 on pred.participant = part.id
														   join users as uses
														     on pred."user" = uses.id
														   join matches as matc
															 on part."match" = matc.id
														   where pred.participant != 3
														     and uses.id = users.id
															 and ((matc.event = $1) or ($1 = -1))
															 and ((pred."user" = $2) or ($2 = -1)))),2) as percent
	from public.predictions as preds
	join public.participants as parts
	  on preds.participant = parts.id
	join public.users       as users
	  on preds."user" = users.id
	join public.matches     as maths
	  on parts."match" = maths.id
	where parts.competitor = maths.winner
	  and ((maths.event = $1) or ($1 = -1))
	  and ((preds."user" = $2) or ($2 = -1))
	group by users.id
	order by count(preds.id) desc

 $function$
