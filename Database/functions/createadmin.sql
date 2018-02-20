CREATE OR REPLACE FUNCTION public.createadmin(id bytea, username text, password text)
 RETURNS void
 LANGUAGE plpgsql
AS $function$

	declare 
		salt text;
		pw text;

	begin	
		salt := gen_salt('bf');
		pw := crypt($3, salt);

		insert into public.auth(id, username, password, salt)
		values(id, username, pw, salt);
    
end
$function$
