-- drop function deleteevent_sp(int2)
CREATE OR REPLACE FUNCTION deleteevent_sp(int2) RETURNS VOID AS
$$
declare idParam int2 = $1;
BEGIN
    delete from public.events
    where id = idParam;
END
$$
  LANGUAGE 'plpgsql';