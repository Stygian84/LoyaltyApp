CREATE TYPE "promo_type_enum" AS ENUM (
  'onetime',
  'ongoing'
);

CREATE TYPE "earn_rate_type_enum" AS ENUM (
  'add',
  'mul'
);

CREATE TYPE "transaction_status_enum" AS ENUM (
  'created',
  'pending',
  'approved',
  'rejected'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar,
  "credit_balance" float NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "contact" int UNIQUE,
  "password" varchar NOT NULL,
  "user_name" varchar NOT NULL UNIQUE,
  "card_tier" int,
  "created_at" timestamp
);

CREATE TABLE "card_tier" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL UNIQUE,
  "tier" int NOT NULL
);

CREATE TABLE "loyalty_program" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "currency_name" varchar NOT NULL,
  "processing_time" varchar NOT NULL,
  "description" varchar,
  "enrollment_link" varchar NOT NULL,
  "terms_condition_link" varchar NOT NULL,
  "format_regex" varchar NOT NULL,
  "partner_code" varchar NOT NULL,
  "initial_earn_rate" float NOT NULL
);

CREATE TABLE "loyalty_membership" (
  "id" bigserial PRIMARY KEY,
  "program" int NOT NULL,
  "name" varchar NOT NULL UNIQUE
);

CREATE TABLE "promotion" (
  "id" bigserial PRIMARY KEY,
  "program" int NOT NULL,
  "promo_type" promo_type_enum NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "earn_rate_type" earn_rate_type_enum NOT NULL,
  "constant" float NOT NULL,
  "card_tier" int,
  "loyalty_membership" int
);

CREATE TABLE "credit_request" (
  "reference_number" bigserial PRIMARY KEY,
  "user_id" int NOT NULL,
  "program" int NOT NULL,
  "member_id" varchar NOT NULL,
  "transaction_time" timestamp DEFAULT (now()),
  "amount" float NOT NULL,
  "transaction_status" transaction_status_enum
);


CREATE INDEX ON "users" ("user_name");

CREATE INDEX ON "users" ("card_tier");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("contact");

CREATE INDEX ON "card_tier" ("name");

CREATE INDEX ON "loyalty_program" ("name");

CREATE INDEX ON "loyalty_membership" ("program");

CREATE INDEX ON "promotion" ("program");

CREATE INDEX ON "credit_request" ("user_id");

CREATE INDEX ON "credit_request" ("program");


ALTER TABLE "users" ADD FOREIGN KEY ("card_tier") REFERENCES "card_tier" ("id");

ALTER TABLE "loyalty_membership" ADD FOREIGN KEY ("program") REFERENCES "loyalty_program" ("id");

ALTER TABLE "promotion" ADD FOREIGN KEY ("program") REFERENCES "loyalty_program" ("id");

ALTER TABLE "promotion" ADD FOREIGN KEY ("card_tier") REFERENCES "card_tier" ("id");

ALTER TABLE "promotion" ADD FOREIGN KEY ("loyalty_membership") REFERENCES "loyalty_membership" ("id");

ALTER TABLE "credit_request" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "credit_request" ADD FOREIGN KEY ("program") REFERENCES "loyalty_program" ("id");

