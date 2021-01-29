-- Table books


CREATE TABLE public.books (
	"id" bigserial PRIMARY KEY,
    "user_id" int8 NOT NULL,
	"title" varchar(255) NOT NULL,
	"pages" integer NOT NULL,
	"created_at" timestamp NOT NULL DEFAULT current_timestamp,
    "updated_at" timestamp NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES public.users("id") ON DELETE NO ACTION
);