CREATE OR REPLACE FUNCTION public.insert_button_style_sp(id uuid, color varchar(18), txt varchar(50), link text, hiding boolean)
 RETURNS TABLE(returnid uuid)
 LANGUAGE plpgsql
AS $function$
begin
	-- insert
		insert into public.button_styles(button_id, button_color, button_text, button_link, is_hiding)
		values (id, color, txt, link, hiding) on conflict(button_id) 
		do update
		set
		button_id = id,
		button_color = color,
		button_text = txt,
		button_link = link,
		is_hiding = hiding;
	return query
	select s.button_id as returnid from public.button_styles s where s.button_id = id;
end $function$;