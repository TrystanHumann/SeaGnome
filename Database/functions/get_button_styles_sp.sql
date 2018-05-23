CREATE OR REPLACE FUNCTION public.get_button_styles_sp()
 RETURNS TABLE(button_id uuid, button_color varchar(18), button_text varchar(50), button_link text)
 LANGUAGE sql
AS $function$
	select button_id, button_color, button_text, button_link
	from public.button_styles
$function$;
