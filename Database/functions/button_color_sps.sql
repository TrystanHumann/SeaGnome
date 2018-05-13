-- Fetch button colors
CREATE OR REPLACE FUNCTION public.get_button_colors(uuid = null)
 RETURNS TABLE(b_guid uuid, hex_code varchar)
 LANGUAGE plpgsql
AS $function$
declare guid uuid = $1;
BEGIN
    if guid is null then
    	return query 
	    select c.button_guid, c.hex_color from public.button_colors c;
    else
	    return query
	    select c.button_guid, c.hex_color from public.button_colors c where c.button_guid = guid;
    end if;
END
$function$;

-- Update button colors
CREATE OR REPLACE FUNCTION public.update_button_color(uuid = null, varchar = null)
 RETURNS TABLE(b_guid uuid, hex_code varchar)
 LANGUAGE plpgsql
AS $function$
declare guid uuid = $1;
declare hex_code varchar = $2;
BEGIN
	insert into public.button_colors(button_guid, hex_color)
	values (guid, hex_code) on conflict(button_guid) 
	do update
	set
	button_guid = guid,
	hex_color = hex_code;	

	-- returning back inserted object
    return query
    select c.button_guid, c.hex_color from public.button_colors c where c.button_guid = guid;
END
$function$;

