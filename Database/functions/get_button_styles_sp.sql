-- drop FUNCTION public.get_button_styles_sp();
CREATE OR REPLACE FUNCTION public.get_button_styles_sp()
 RETURNS TABLE(button_id uuid, button_color varchar(18), button_text varchar(50), button_link text, is_hiding boolean)
 LANGUAGE sql
AS $function$
	select button_id, button_color, button_text, button_link, is_hiding
	from public.button_styles
$function$;
