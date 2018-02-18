-- drop function public.insert_game_sp(text)
CREATE OR REPLACE FUNCTION public.insert_game_sp(gamename text) returns table(returnid int)
 LANGUAGE plpgsql
AS $$
declare fetchID int;
begin
	-- Assuming they don't update competitors
	fetchID := (select g.id from public.games g where g."name" = gamename);
	if fetchID is null then
	-- insert
		insert into public.games(name)
		values (gamename);
	end if;
	return query
	select g.id as returnid from public.games g where g."name" = gamename;
	--RAISE INFO 'Out variable: %', outUserID;
end $$;


