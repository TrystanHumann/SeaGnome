CREATE OR REPLACE FUNCTION public.insert_prediction_sp
(userid integer, participantid integer)
 RETURNS TABLE
(returnid integer)
 LANGUAGE plpgsql
AS $function$
begin
    -- insert
    insert into public.predictions
        ("user", participant)
    values
        (userid, participantid)
    on conflict
    ("user", participant) 
		do
    update
		set
		"user" = userid,
		participant = participantid;
    return query
    select p.id as returnid
    from public.predictions p
    where p."user" = userid and p.participant = participantid;
--RAISE INFO 'Out variable: %', outUserID;
end
$function$
