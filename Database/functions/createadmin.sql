CREATE OR REPLACE FUNCTION public.createadmin(id bytea, username text, password text)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$

	declare 
		salt text;
		pw text;

	begin	
		salt := gen_salt('crypt-bf/5');
		pw := crypt($3, salt);

		insert into public.auth(id, username, password, salt)
		values(id, username, pw, salt);
    
end
$function$
