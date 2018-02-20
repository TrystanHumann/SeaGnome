<<<<<<< HEAD
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
=======
-- drop function public.insert_prediction_sp(int, int)
CREATE OR REPLACE FUNCTION public.insert_prediction_sp(userid int, participantid int) returns table(returnid int)
 LANGUAGE plpgsql
AS $$
declare fetchID int;
begin
	-- Assuming they don't update competitors
	fetchID := (select p.id from public.predictions p where p."user" = userid and p.participant = participantid);
	if fetchID is null then
	-- insert
		insert into public.predictions("user", participant)
		values (userid, participantid);
	else 
		-- update 
		update public.predictions
		set participant = participantid
		where user = userid;
	end if;
	return query
	select p.id as returnid from public.predictions p where p."user" = userid and p.participant = participantid;	
	--RAISE INFO 'Out variable: %', outUserID;
end $$;


>>>>>>> d305009f3ebb93b3f7e7c23c87f02ba6be5d6fd7
