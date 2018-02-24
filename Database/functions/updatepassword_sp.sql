CREATE OR REPLACE FUNCTION public.updatepassword_sp(username text, oldpassword text, newpassword text)
 RETURNS TABLE(id bytea, username text, token bytea, expires timestamp with time zone)
 LANGUAGE sql
AS $function$

	update public.auth
	set password = crypt(newPassword, public.auth.salt)
	where public.auth.username = username
      and public.auth.password = crypt(oldPassword, public.auth.salt)
    returning public.auth.id, public.auth.username, public.auth.token, public.auth.expires;
$function$;
