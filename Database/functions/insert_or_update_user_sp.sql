-- drop function insert_or_update_user_sp(text, text)
CREATE OR REPLACE FUNCTION public.insert_or_update_user_sp(twitchUsername text, twitter text = null) returns table(returnid int)
 LANGUAGE plpgsql
AS $$
declare fetchID int;
begin
	fetchID := (select u.id from public.users u where u.twitchun = twitchUsername);
	if fetchID is null then
	-- insert
		insert into public.users(twitchun, twitterhandle)
		values (twitchUsername, twitter);
	else 
	-- update
		update public.users
		set twitterhandle = twitter
		where id = fetchID; 
	end if;
	return query
	select u.id as returnid from public.users u where u.twitchun = twitchUsername;
	--RAISE INFO 'Out variable: %', outUserID;
end $$;


