CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_product_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
	"success" boolean not NULL DEFAULT false
);

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_product_id") REFERENCES "product" ("id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_product_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_product_id");

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';