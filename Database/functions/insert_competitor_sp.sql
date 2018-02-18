-- drop function public.insert_competitor_sp(text)
CREATE OR REPLACE FUNCTION public.insert_competitor_sp(competitorname text) returns table(returnid smallint)
 LANGUAGE plpgsql
AS $$
declare fetchID int;
begin
	-- Assuming they don't update competitors
	fetchID := (select c.id from public.competitors c where c."name" = competitorname);
	if fetchID is null then
	-- insert
		insert into public.competitors(name)
		values (competitorname);
	end if;
	return query
	select c.id as returnid from public.competitors c where c."name" = competitorname;
	--RAISE INFO 'Out variable: %', outUserID;
end $$;


