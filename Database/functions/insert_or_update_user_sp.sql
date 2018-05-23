CREATE OR REPLACE FUNCTION public.insert_or_update_user_sp(twitchusername text, twitter text DEFAULT NULL::text)
 RETURNS TABLE(returnid integer)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
	-- insert
		insert into public.users(twitchun, twitterhandle)
		values (twitchUsername, twitter) on conflict(twitchun) 
		do update
		set
		twitchun = twitchusername,
		twitterhandle = twitter;
	return query
	select u.id as returnid from public.users u where u.twitchun = twitchUsername;
	--RAISE INFO 'Out variable: %', outUserID;
end $function$;
