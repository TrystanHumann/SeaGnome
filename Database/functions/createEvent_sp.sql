-- drop function createevent_sp(name text)
CREATE OR REPLACE FUNCTION createevent_sp(name text) RETURNS VOID AS
$$
declare nameParam text = $1;
BEGIN
    INSERT INTO public.events (name, complete, created_date) VALUES (nameParam, false, timezone('utc', now()));
END
$$
  LANGUAGE 'plpgsql';
-- CREATE OR REPLACE FUNCTION createEvent_sp(name text) RETURNS VOID AS
-- $$
-- declare nameParam text = $1;
-- BEGIN
--     INSERT INTO public.events (name, complete, created_date) VALUES (nameParam, false, CAST(current_timestamp::timestamp at time zone 'America/New_York' AS timestamp with time zone));
-- END
-- $$
--   LANGUAGE 'plpgsql';