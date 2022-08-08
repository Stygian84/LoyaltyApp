DELETE FROM "credit_request";
DELETE FROM "users";


ALTER TABLE "users" 
ALTER COLUMN "full_name" SET NOT NULL;

ALTER TABLE "users" 
ALTER COLUMN "full_name" SET DEFAULT '';

ALTER TABLE "users" 
ALTER COLUMN "contact" SET NOT NULL;


ALTER TABLE "users" 
ALTER COLUMN "card_tier" SET NOT NULL;

ALTER TABLE "users" 
ALTER COLUMN "card_tier" SET DEFAULT 0;

ALTER TABLE "users" 
ALTER COLUMN "created_at" SET DEFAULT (now())


