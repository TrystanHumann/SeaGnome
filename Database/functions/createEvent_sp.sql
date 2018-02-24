CREATE OR REPLACE FUNCTION public.createevent_sp(name text)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
declare nameParam text = $1;
BEGIN
    INSERT INTO public.events (name, complete, created_date, active) VALUES (nameParam, false, timezone('utc', now()), false);
END
$function$
