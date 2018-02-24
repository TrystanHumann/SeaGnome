CREATE OR REPLACE FUNCTION public.getUserId_sp
("user" text)
 RETURNS TABLE
(id smallint, twitch text, twitter text)
 LANGUAGE plpgsql
AS $function$
begin
    return query
    select u.id as "id", u.twitchun as "twitch", u.twitterhandle as "twitter"
    from public.users u
    where u.twitchun = "user";
end;
$function$
