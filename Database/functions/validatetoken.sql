CREATE OR REPLACE FUNCTION public.validatetoken(id bytea, token bytea)
 RETURNS TABLE(valid boolean)
 LANGUAGE sql
AS $function$

	select case
			when public.auth.token = $2 then true 
			else false
		   end as valid
	from public.auth
	where public.auth.id = $1
	  and public.auth.expires > now()
	
$function$;
