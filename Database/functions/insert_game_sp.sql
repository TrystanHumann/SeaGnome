CREATE OR REPLACE FUNCTION public.insert_game_sp
(gamename text)
 RETURNS TABLE
(returnid integer)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
    -- insert
    insert into public.games
        ("name")
    values
        (gamename)
    on conflict
    ("name") 
		do
    update
		set
		"name" = gamename;
    return query
    select g.id as returnid
    from public.games g
    where g."name" = gamename;
--RAISE INFO 'Out variable: %', outUserID;
end
$function$
