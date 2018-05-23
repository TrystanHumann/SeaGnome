CREATE OR REPLACE FUNCTION public.getuserscore(event integer, usr integer)
 RETURNS TABLE("user" text, total bigint, percent numeric)
 LANGUAGE sql
AS $function$

select "count".un as "user"
     , sum("count".correct) as total
     , case
         when sum("count".nonAbstained) = 0 then 0.00
         else round(100.00 * 
                     cast(sum("count".correct) as numeric) /
                     cast(sum("count".nonAbstained) as numeric),2)
       end as percent
from(
select us.twitchun as un
     , case 
         when co.id = ma.winner
           then 1
         else 0 
       end as correct
     , case
         when ma.winner is not null and co.id != 3
           then 1
         else 0
       end as nonAbstained
from public.predictions as pr
join public.participants as pa
  on pr.participant = pa.id
join public.users as us
  on pr."user" = us.id
join public.competitors as co
  on pa.competitor = co.id
join public.matches as ma
  on pa."match" = ma.id
 and ((ma.event = $1) or ($1 = -1))
where ((pr."user" = $2) or ($2 = -1))
) as "count"
group by "count".un
order by total desc, percent desc, "count".un asc

 $function$