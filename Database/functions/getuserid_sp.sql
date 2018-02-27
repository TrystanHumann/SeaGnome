CREATE OR REPLACE FUNCTION public.getuserid_sp
("user" text)
 RETURNS TABLE
(id integer, twitch text, twitter text)
 LANGUAGE plpgsql
AS $function$
begin
        return query
        select u.id as "id", u.twitchun as "twitch", u.twitterhandle as "twitter"
        from public.users u
        where lower(u.twitchun) = lower("user");
end;
$function$
