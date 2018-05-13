---------------- TABLES ----------------

-- EVENTS
drop table if exists public.events;
CREATE TABLE public.events (
	id smallserial NOT NULL,
	"name" text NOT NULL,
	complete bool NOT NULL,
	created_date timestamptz NOT NULL,
	active bool NOT NULL,
	CONSTRAINT events_pkey PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX events_pk_index ON events USING btree (id) ;

-- GAMES
drop table if exists public.games;
CREATE TABLE public.games (
	id serial NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT games_pk PRIMARY KEY (id),
	CONSTRAINT games_unique_game_name UNIQUE (name)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX games_pk_index ON games USING btree (id) ;

-- COMPETITORS
drop table if exists public.competitors;
CREATE TABLE public.competitors (
	id smallserial NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT competitors_pk PRIMARY KEY (id),
	CONSTRAINT competitors_unique_competitors UNIQUE (name)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX competitors_pk_index ON competitors USING btree (id) ;

-- AUTH
drop table if exists public.auth;
CREATE TABLE public.auth (
	id bytea NOT NULL,
	username text NOT NULL,
	password text NOT NULL,
	salt text NOT NULL,
	token bytea NULL,
	expires timestamptz NULL,
	CONSTRAINT auth_pk PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;

-- MATCHES
drop table if exists public.matches;
CREATE TABLE public.matches (
	id serial NOT NULL,
	event int4 NOT NULL,
	game int4 NOT NULL,
	scheduled timestamptz NULL,
	winner int4 NULL,
	CONSTRAINT matches_pkey PRIMARY KEY (id),
	CONSTRAINT matches_unique_event_game UNIQUE (event, game),
	CONSTRAINT matches_game_games_id_fk FOREIGN KEY (game) REFERENCES games(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT matches_winner_competitors_id_fk FOREIGN KEY (winner) REFERENCES competitors(id) ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
	OIDS=FALSE
) ;
CREATE INDEX fki_matches_event_events_id_fk ON matches USING btree (event) ;
CREATE INDEX fki_matches_game_games_id_fk ON matches USING btree (game) ;
CREATE INDEX fki_matches_winner_competitors_id_fk ON matches USING btree (winner) ;
CREATE UNIQUE INDEX matches_pk_index ON matches USING btree (id) ;

-- PARTICIPANTS
drop table if exists public.participants;
CREATE TABLE public.participants (
	id serial NOT NULL,
	"match" int4 NOT NULL,
	competitor int4 NOT NULL,
	CONSTRAINT participants_pkey PRIMARY KEY (id),
	CONSTRAINT participants_unique_competitor_match UNIQUE (match, competitor),
	CONSTRAINT particiants_match_match_id_fk FOREIGN KEY (match) REFERENCES matches(id),
	CONSTRAINT participant_competitor_competitors_id_fk FOREIGN KEY (competitor) REFERENCES competitors(id)
)
WITH (
	OIDS=FALSE
) ;
CREATE INDEX fki_particiants_match_match_id_fk ON participants USING btree (match) ;
CREATE INDEX fki_participant_competitor_competitors_id_fk ON participants USING btree (competitor) ;
CREATE UNIQUE INDEX participants_pk_index ON participants USING btree (id) ;

-- USERS
drop table if exists public.users;
CREATE TABLE public.users (
	id serial NOT NULL,
	twitchun text NOT NULL,
	twitterhandle text NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique_twitchun UNIQUE (twitchun)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX users_pk_index ON users USING btree (id) ;

-- PREDICTIONS
drop table if exists public.predictions;
CREATE TABLE public.predictions (
	id serial NOT NULL,
	"user" int4 NOT NULL,
	participant int4 NOT NULL,
	CONSTRAINT predictions_pkey PRIMARY KEY (id),
	CONSTRAINT predictions_unique_user_participant UNIQUE ("user", participant),
	CONSTRAINT predictions_participant_participants_id_fk FOREIGN KEY (participant) REFERENCES participants(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT predictions_user_users_id_fk FOREIGN KEY ("user") REFERENCES users(id) ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
	OIDS=FALSE
) ;
CREATE INDEX fki_predictions_participant_participants_id_fk ON predictions USING btree (participant) ;
CREATE INDEX fki_predictions_user_users_id_fk ON predictions USING btree ("user") ;
CREATE UNIQUE INDEX predictions_id_pk ON predictions USING btree (id) ;

-- BONUS QUESTIONS
drop table if exists public.bonusquestions;
CREATE TABLE public.bonusquestions (
	id serial NOT NULL,
	question text NOT NULL,
	event int4 NOT NULL,
	tiebreaker bool NOT NULL,
	CONSTRAINT bonusquestions_pkey PRIMARY KEY (id),
	CONSTRAINT bonusquestions_unique_question_event UNIQUE (question, event)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX bonusquestions_pk_index ON bonusquestions USING btree (id) ;
CREATE INDEX fki_bonusquestions_event_events_id_fk ON bonusquestions USING btree (event) ;

-- BONUS ANSWERS
drop table if exists public.bonusanswers;
CREATE TABLE public.bonusanswers (
	id serial NOT NULL,
	question int4 NOT NULL,
	answer text NOT NULL,
	CONSTRAINT bonusanswers_pkey PRIMARY KEY (id),
	CONSTRAINT bonusanswers_question_fkey FOREIGN KEY (question) REFERENCES bonusquestions(id) ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX bonusanswers_pk_index ON bonusanswers USING btree (id) ;
CREATE INDEX fki_bonusanswers_question_to_bonusquestions_fk ON bonusanswers USING btree (question) ;

-- STREAMERS
drop table if exists public.streamers;
CREATE TABLE public.streamers (
	id serial NOT NULL,
	tag text NOT NULL,
	active bool NOT NULL DEFAULT true,
	CONSTRAINT streamers_pk PRIMARY KEY (id),
	CONSTRAINT streamers_unique_streamer UNIQUE (tag)
)
WITH (
	OIDS=FALSE
) ;
CREATE UNIQUE INDEX streamers_id_idx ON streamers USING btree (id) ;

-- BUTTON COLOR
drop table if exists public.button_colors;
CREATE TABLE "public".button_colors (
	button_guid uuid NOT NULL,
	hex_color varchar(7) NOT NULL
)
WITH (
	OIDS=FALSE
) ;
COMMENT ON TABLE "public".button_colors IS 'Table used to store button color' ;


---------------- FUNCTIONS ----------------

-- CREATE ADMIN
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

-- CREATE EVENT
CREATE OR REPLACE FUNCTION public.createevent_sp(name text)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
declare nameParam text = $1;
BEGIN
    INSERT INTO public.events (name, complete, created_date, active) VALUES (nameParam, false, timezone('utc', now()), false);
END
$function$;

-- DELETE EVENT
CREATE OR REPLACE FUNCTION public.deleteevent_sp(smallint)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
declare idParam int2 = $1;
BEGIN
    delete from public.events
    where id = idParam;
END
$function$;

-- GET COMPETITOR ID
CREATE OR REPLACE FUNCTION public.get_competitorid_sp(competitor text)
 RETURNS TABLE(returnid smallint)
 LANGUAGE plpgsql
AS $function$
begin
	
	return query
	select c.id as returnid from public.competitors c where c."name" = competitor;

end $function$;

-- GET GAME ID
CREATE OR REPLACE FUNCTION public.get_gameid_sp(gamename text)
 RETURNS TABLE(returnid integer)
 LANGUAGE plpgsql
AS $function$
begin
	return query
	select g.id as returnid from public.games g where g."name" = gamename;
end $function$;

-- GET ACTIVE EVENT
CREATE OR REPLACE FUNCTION public.getactiveevent_sp()
 RETURNS TABLE(id smallint, name text, complete boolean, created_date timestamp with time zone, activeflag boolean)
 LANGUAGE plpgsql
AS $function$
	begin
        return query 
        select e.id as "id", e.name as "name", e.complete as "complete", e.created_date as "created_date", e.active as "activeflag"
        from public.events e 
        where e.active = true;
	end;
$function$;

-- GET ACTIVE STREAMERS
CREATE OR REPLACE FUNCTION public.getactivestreamers(max integer)
 RETURNS TABLE(id integer, tag text, active boolean)
 LANGUAGE sql
AS $function$

select id, tag, active
from public.streamers
where active = true
limit $1

 $function$;

-- GET COMPETITORS BY MATCH
CREATE OR REPLACE FUNCTION public.getcompetitorsbymatch(matchsearch integer)
 RETURNS TABLE(id smallint, compname text)
 LANGUAGE sql
AS $function$

	select c.id     as id
	     , c."name" as compname
	from public.participants as p
	join public.competitors as c
	  on p.competitor = c.id
	where p."match" = $1;

$function$;

-- GET EVENT RESULTS
CREATE OR REPLACE FUNCTION public.geteventresults(event integer)
 RETURNS TABLE(name text, wins bigint, matches bigint)
 LANGUAGE sql
AS $function$

	select com."name"
	     , count(par.id) as wins
	     , (select count(p2.id) 
	          from public.participants as p2
	    	  join public.matches      as m2
			    on p2."match" = m2.id
			   and m2.winner is not null
			   and ((m2.event = $1) or ($1 = -1))
			  join public.competitors  as c2
			    on p2.competitor = c2.id
			   and c2."name" = com."name") as matches
	from public.participants as par
	join public.matches      as mat
	  on par."match" = mat.id
	 and mat.winner = par.competitor
	 and ((mat.event = $1) or ($1 = -1))
	join public.competitors  as com
	  on par.competitor = com.id
	group by com."name"

$function$;

-- GET EVENTS
CREATE OR REPLACE FUNCTION public.getevents_sp()
 RETURNS TABLE(id smallint, name text, complete boolean, created_date timestamp with time zone)
 LANGUAGE plpgsql
AS $function$
	begin
        return query 
        select e.id as "id", e.name as "name", e.complete as "complete", e.created_date as "created_date"
        from public.events e 
        order by e.created_date desc;
	end;
$function$;

-- GET FUTURE PREDICTIONS
CREATE OR REPLACE FUNCTION public.getfuturepredictions(event integer)
 RETURNS TABLE(game text, participant text, votes bigint)
 LANGUAGE sql
AS $function$
	select trim(gam."name") as game, trim(com."name") as participant, count(com.id) as votes
	from public.predictions as pre
	join public.users as use
	  on pre."user" = use.id
	join public.participants as par
	  on pre.participant = par.id
	join public.competitors as com
	  on par.competitor = com.id
	join public.matches as mat
	  on par."match" = mat.id
	join public.games as gam
	  on mat.game = gam.id
	where mat.scheduled > now()
	  and ((mat.event = $1) or $1 = -1)
	group by com.id, gam."name", mat.scheduled
	order by mat.scheduled asc, com.id desc
$function$;

-- GET GAME PREDICTIONS
CREATE OR REPLACE FUNCTION public.getgamepredictions(event integer, "user" integer)
 RETURNS TABLE(game text, prediction text, winner text)
 LANGUAGE sql
AS $function$
	select gam."name" as game
	     , com."name" as prediction
	     , coalesce((select "name" from public.competitors where id = mat.winner), '') as winner
	from public.predictions  as pre
	join public.participants as par
	  on pre.participant = par.id
	join public.competitors  as com
	  on par.competitor = com.id
	join public.matches      as mat
	  on par."match" = mat.id
	 and ((mat.event = $1) or ($1 = -1))
	join public.games        as gam
	  on mat.game = gam.id
	where ((pre."user" = $2) or ($2 = -1))
	order by mat.scheduled asc
$function$;

-- GET MATCHES BY EVENT
CREATE OR REPLACE FUNCTION public.getmatchesbyevent(event integer)
 RETURNS TABLE(id integer, game text)
 LANGUAGE sql
AS $function$

	select m.id     as id
	     , g."name" as game
	from public.matches as m
	join public.games as g
	  on m.game = g.id
	where m.event = $1;

$function$;

-- GET USER ID
CREATE OR REPLACE FUNCTION public.getuserid_sp("user" text)
 RETURNS TABLE(id integer, twitch text, twitter text)
 LANGUAGE plpgsql
AS $function$
	begin
        return query 
        select u.id as "id", u.twitchun as "twitch", u.twitterhandle as "twitter"
        from public.users u
        where u.twitchun = "user";
	end;
$function$;

-- GET USER SCORE
CREATE OR REPLACE FUNCTION public.getuserscore(event integer, usr integer)
 RETURNS TABLE("user" text, total bigint, percent numeric)
 LANGUAGE sql
AS $function$

	select users.twitchun
	     , count(preds.id) as correct
	     , round(100 * (cast (count(preds.id) as numeric)/(select count(matc.id)
			                                               from predictions as pred
														   join participants as part
															 on pred.participant = part.id
														   join users as uses
														     on pred."user" = uses.id
														   join matches as matc
															 on part."match" = matc.id
														   where pred.participant != 29
														     and uses.id = users.id
															 and ((matc.event = $1) or ($1 = -1))
															 and ((pred."user" = $2) or ($2 = -1)))),2) as percent
	from public.predictions as preds
	join public.participants as parts
	  on preds.participant = parts.id
	join public.users       as users
	  on preds."user" = users.id
	join public.matches     as maths
	  on parts."match" = maths.id
	where parts.competitor = maths.winner
	  and ((maths.event = $1) or ($1 = -1))
	  and ((preds."user" = $2) or ($2 = -1))
	group by users.id
	order by count(preds.id) desc

 $function$;

-- INSERT COMPETITOR
CREATE OR REPLACE FUNCTION public.insert_competitor_sp(competitorname text)
 RETURNS TABLE(returnid smallint)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
	-- Assuming they don't update competitors
	fetchID := (select c.id from public.competitors c where c."name" = competitorname);
	if fetchID is null then
	-- insert
		insert into public.competitors(name)
		values (competitorname);
	end if;
	return query
	select c.id as returnid from public.competitors c where c."name" = competitorname;
	--RAISE INFO 'Out variable: %', outUserID;
end $function$;

-- INSERT GAME
CREATE OR REPLACE FUNCTION public.insert_game_sp(gamename text)
 RETURNS TABLE(returnid integer)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
	-- insert
		insert into public.games("name")
		values (gamename) on conflict("name") 
		do update
		set
		"name" = gamename;
	return query
	select g.id as returnid from public.games g where g."name" = gamename;
	--RAISE INFO 'Out variable: %', outUserID;
end $function$;

-- INSERT MATCH
CREATE OR REPLACE FUNCTION public.insert_match_sp(ev integer, ga integer, sch timestamp without time zone DEFAULT NULL::timestamp without time zone, win integer DEFAULT NULL::integer)
 RETURNS TABLE(returnid integer)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
	-- insert
	insert into public.matches(event, game, scheduled, winner)
	values (ev, ga, sch, win)
	on conflict(event, game) do update 
	set
		event = ev,
		game = ga, 
		scheduled = coalesce(sch, (select scheduled from public.matches where event = ev and game = ga)),
		winner = coalesce(win, (select winner from public.matches where event = ev and game = ga));
	return query
	select m.id as returnid from public.matches m where m.event = ev and m.game = ga;	
	--RAISE INFO 'Out variable: %', outUserID;
end $function$;

-- INSERT OR UPDATE USER
CREATE OR REPLACE FUNCTION public.insert_or_update_user_sp(twitchusername text, twitter text DEFAULT NULL::text)
 RETURNS TABLE(returnid integer)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
	-- insert
		insert into public.users(twitchun, twitterhandle)
		values (twitchUsername, twitter) on conflict(twitchun) 
		do update
		set
		twitchun = twitchusername,
		twitterhandle = twitter;
	return query
	select u.id as returnid from public.users u where u.twitchun = twitchUsername;
	--RAISE INFO 'Out variable: %', outUserID;
end $function$;

-- INSERT PARTICIPANT
CREATE OR REPLACE FUNCTION public.insert_participant_sp(matchid integer, competitorid integer)
 RETURNS TABLE(returnid integer)
 LANGUAGE plpgsql
AS $function$
declare fetchID int;
begin
	-- Assuming they don't update competitors
	fetchID := (select p.id from public.participants p where p."match" = matchid and p.competitor = competitorid);
	if fetchID is null then
	-- insert
		insert into public.participants(match, competitor)
		values (matchid, competitorid);
	end if;
	return query
	select p.id as returnid from public.participants p where p."match" = matchid and p.competitor = competitorid;
	--RAISE INFO 'Out variable: %', outUserID;
end $function$;

-- INSERT PREDICTION
CREATE OR REPLACE FUNCTION public.insert_prediction_sp(userid integer, participantid integer)
 RETURNS TABLE(returnid integer)
 LANGUAGE plpgsql
AS $function$
begin
	-- insert
		insert into public.predictions("user", participant)
		values (userid, participantid) on conflict("user", participant) 
		do update
		set
		"user" = userid,
		participant = participantid;
	return query
	select p.id as returnid from public.predictions p where p."user" = userid and p.participant = participantid;	
	--RAISE INFO 'Out variable: %', outUserID;
end $function$;

-- INSERT STREAMER
CREATE OR REPLACE FUNCTION public.insertstreamer
(streamerone text, streamertwo text)
 RETURNS void
 LANGUAGE sql
AS $function$

truncate public.streamers;

insert into public.streamers
    (tag, active)
values($1, true);

insert into public.streamers
    (tag, active)
values($2, true);


$function$
-- LOGIN
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

-- UPDATE EVENT ACTIVE
CREATE OR REPLACE FUNCTION public.updateevent_active_sp(integer, boolean)
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
$function$;

-- UPDATE EVENT
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
$function$;

-- UPDATE PASSWORD
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

-- UPDATE STREAMER
CREATE OR REPLACE FUNCTION public.updatestreamer(id integer, active boolean)
 RETURNS text
 LANGUAGE sql
AS $function$

update public.streamers
set active = $2
where id = $1
returning tag

 $function$;

 -- VALIDATE TOKEN
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

-- Fetch button colors
CREATE OR REPLACE FUNCTION public.get_button_color(uuid = null)
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
	values (guid, hex_code) on conflict(guid) 
	do update
	set
	button_guid = guid,
	hex_color = hex_code;	

	-- returning back inserted object
    return query
    select c.button_guid, c.hex_color from public.button_colors c where c.button_guid = guid;
END
$function$;

