CREATE TABLE public.users (
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL,
	short_url varchar NOT NULL,
	long_url varchar NULL,
);