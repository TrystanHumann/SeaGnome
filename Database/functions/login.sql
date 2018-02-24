CREATE OR REPLACE FUNCTION public.login(username text, password text, token bytea)
 RETURNS TABLE(id bytea, username text, token bytea, expires timestamp with time zone)
 LANGUAGE sql
AS $function$

	update public.auth
	set token = $3, expires = now() + interval '1 day'
	where public.auth.username = $1
      and public.auth.password = crypt($2, public.auth.salt)
    returning public.auth.id, public.auth.username, public.auth.token, public.auth.expires;

$function$;
