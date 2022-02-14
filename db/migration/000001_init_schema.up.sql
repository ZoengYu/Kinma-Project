CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "content" text NOT NULL,
  "product_tag" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_update" timestamptz
);

CREATE TABLE "fundraise" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "target_amount" bigint NOT NULL,
  "progress_amount" bigint NOT NULL,
  "success" boolean,
  "start_date" timestamptz NOT NULL DEFAULT (now()),
  "end_date" timestamptz
);

ALTER TABLE "product" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "fundraise" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "product" ("account_id");

CREATE INDEX ON "product" ("title");

CREATE INDEX ON "product" ("product_tag");

CREATE INDEX ON "fundraise" ("success");

CREATE INDEX ON "fundraise" ("progress_amount");

CREATE INDEX ON "fundraise" ("success", "target_amount");

COMMENT ON COLUMN "fundraise"."target_amount" IS 'must be positive';

COMMENT ON COLUMN "fundraise"."progress_amount" IS 'must be positive';
