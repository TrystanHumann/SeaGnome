-- drop function updateevent_sp(int2, boolean)
CREATE OR REPLACE FUNCTION updateevent_sp(int2, boolean) RETURNS VOID AS
$$
declare 
	idParam int2 = $1;
	compParam boolean = $2;
BEGIN
    update public.events
    set complete = compParam
    where id = idParam;
END
$$
  LANGUAGE 'plpgsql';

-- select public".updateevent_sp(9::int2, true::boolean)

