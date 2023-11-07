CREATE TABLE "users"
(
    "id" bigserial NOT NULL PRIMARY KEY,
    "name" character varying(255) NOT NULL UNIQUE,
    "password" text NOT NULL,
    "email" character varying(255) NOT NULL UNIQUE,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone
);