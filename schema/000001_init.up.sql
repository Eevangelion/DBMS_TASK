CREATE TABLE "Favorite jokes"
(
	joke_id integer NOT NULL,
	user_id integer NOT NULL,
	CONSTRAINT "Favorite jokes_pkey" PRIMARY KEY (joke_id, user_id),
	FOREIGN KEY (joke_id) REFERENCES public."Jokes" (id) MATCH SIMPLE 
	ON UPDATE NO ACTION
	ON DELETE CASCADE
	NOT VALID,
	FOREIGN KEY (user_id)
	REFERENCES public."Users" (id) MATCH SIMPLE
	ON UPDATE NO ACTION
	ON DELETE CASCADE
	NOT VALID
);
	
CREATE TABLE "Jokes"
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
	header character varying(128) COLLATE pg_catalog."default",
	description text COLLATE pg_catalog."default" NOT NULL,
	rating integer NOT NULL DEFAULT 0,
	author_id integer NOT NULL,
	creation_date timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT "Jokes_pkey" PRIMARY KEY (id),
	CONSTRAINT rating_check CHECK (rating >= 0),
	FOREIGN KEY (author_id) REFERENCES public."Users" (id) MATCH SIMPLE
	ON UPDATE NO ACTION
	ON DELETE CASCADE
	NOT VALID
);

CREATE TABLE "Reports"
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
	description text COLLATE pg_catalog."default" NOT NULL,
	receiver_joke_id integer NOT NULL,
	sender_id integer NOT NULL,
	receiver_id integer NOT NULL,
	CONSTRAINT "Reports_pkey" PRIMARY KEY (id),
	FOREIGN KEY (receiver_id) REFERENCES public."Users" (id) MATCH SIMPLE
	ON UPDATE NO ACTION
	ON DELETE CASCADE
	NOT VALID,
	FOREIGN KEY (receiver_joke_id) REFERENCES public."Jokes" (id) MATCH SIMPLE
	ON UPDATE NO ACTION
	ON DELETE CASCADE
	NOT VALID,
	FOREIGN KEY (sender_id) REFERENCES public."Users" (id) MATCH SIMPLE
	ON UPDATE NO ACTION
	ON DELETE CASCADE
	NOT VALID
);

CREATE TABLE "Tags"
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
	name character varying(128) COLLATE pg_catalog."default" NOT NULL,
	CONSTRAINT "Tags_pkey" PRIMARY KEY (id),
	CONSTRAINT "Uniq_name" UNIQUE (name)
);

CREATE TABLE "TagsJokes"
(
	tag_id integer NOT NULL,
	joke_id integer NOT NULL,
	CONSTRAINT "TagsJokes_pkey" PRIMARY KEY (tag_id, joke_id),
	FOREIGN KEY (tag_id) REFERENCES public."Tags" (id) MATCH SIMPLE
	ON UPDATE NO ACTION
	ON DELETE CASCADE
    NOT VALID,
	FOREIGN KEY (joke_id) REFERENCES public."Jokes" (id) MATCH SIMPLE
	ON UPDATE NO ACTION
	ON DELETE CASCADE
	NOT VALID

);

CREATE TABLE "Users"
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
	name character varying(64) COLLATE pg_catalog."default" NOT NULL,
	email character varying(128) COLLATE pg_catalog."default" NOT NULL,
	reports integer NOT NULL DEFAULT 0,
	remaining_reports integer NOT NULL DEFAULT 3,
	role character varying(64) COLLATE pg_catalog."default" NOT NULL,
	unban_date date NOT NULL DEFAULT '1861-03-03'::date,
	transformed_password character varying(128) COLLATE pg_catalog."default" NOT NULL,
	CONSTRAINT "Users_pkey" PRIMARY KEY (id),
	CONSTRAINT "Uniq_users_characteristics" UNIQUE (name, email),
	CONSTRAINT "ReportsAreRemaining" CHECK (remaining_reports >= 0) NOT VALID
);

CREATE TABLE "GithubUsers"
(
	git_id integer NOT NULL,
	inner_id integer NOT NULL,
	CONSTRAINT "GithubUsers_pkey" PRIMARY KEY (git_id, inner_id)
);

CREATE TABLE "UserSubscribes"
(
	receiver_id integer NOT NULL,
	sender_id integer NOT NULL,
	PRIMARY KEY (receiver_id, "sender_id"),
	FOREIGN KEY (receiver_id) REFERENCES public."Users" (id) MATCH SIMPLE
	ON UPDATE NO ACTION
	ON DELETE CASCADE
	NOT VALID,
	FOREIGN KEY (sender_id) REFERENCES public."Users" (id) MATCH SIMPLE
	ON UPDATE NO ACTION
	ON DELETE CASCADE
	NOT VALID
);