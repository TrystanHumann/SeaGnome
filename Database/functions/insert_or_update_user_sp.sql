-- drop function insert_or_update_user_sp(text, text, int)
CREATE OR REPLACE FUNCTION public.insert_or_update_user_sp(twitchUsername text, twitter text = null, out outUserID int) returns int
 LANGUAGE plpgsql
AS $$
declare fetchID int;
declare outerUserID int;
begin
	fetchID := (select u.id from public.users u where u.twitchun = twitchUsername);
	if fetchID is null then
	-- insert
		insert into public.users(twitchun, twitterhandle)
		values (twitchUsername, twitter);
		outUserID := (select u.id from public.users u where u.twitchun = twitchUsername);
	else 
	-- update
		update public.users
		set twitterhandle = twitter
		where id = fetchID;
		outUserID := fetchID;
	end if;
	--RAISE INFO 'Out variable: %', outUserID;
end $$;


