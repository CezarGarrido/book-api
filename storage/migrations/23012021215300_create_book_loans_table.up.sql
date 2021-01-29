-- Table book_loans

CREATE TABLE public.book_loans (
	"id" bigserial PRIMARY KEY,
    "book_id" int8 NOT NULL,
    "from_user_id" int8 NOT NULL,
	"to_user_id" int8 NOT NULL,
    "lent_at" timestamp NOT NULL,
    "returned_at" timestamp,
	"created_at" timestamp NOT NULL DEFAULT current_timestamp,
    "updated_at" timestamp NOT NULL,
    FOREIGN KEY ("book_id") REFERENCES books("id")  ON DELETE NO ACTION,
    FOREIGN KEY ("from_user_id") REFERENCES users("id") ON DELETE NO ACTION,
    FOREIGN KEY ("to_user_id") REFERENCES users("id") ON DELETE NO ACTION
);

