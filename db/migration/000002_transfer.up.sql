CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_fundraise_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
	"success" boolean not NULL DEFAULT false
);

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_fundraise_id") REFERENCES "fundraise" ("id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_fundraise_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_fundraise_id");

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';