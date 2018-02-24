<<<<<<< HEAD
CREATE OR REPLACE FUNCTION public.updateevent_active_sp
(integer, boolean)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
declare 
	idParam int = $1;
	activeParam bool = $2;
BEGIN
  update public.events
    set active = false
    where id != idParam;

  update public.events
    set active = activeParam
    where id = idParam;
END
$function$
=======
CREATE OR REPLACE FUNCTION public.updateevent_sp(smallint, boolean)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
declare 
	idParam int2 = $1;
	compParam boolean = $2;
BEGIN
    update public.events
    set complete = compParam
    where id = idParam;
END
$function$
>>>>>>> origin/development
