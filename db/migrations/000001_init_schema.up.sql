CREATE TABLE "User" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Notes" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGSERIAL NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_updated" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ShareHistory" (
  "id" BIGSERIAL PRIMARY KEY,
  "owner_id" BIGSERIAL NOT NULL,
  "shared_id" BIGSERIAL NOT NULL,
  "shared_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "User" ("id");

CREATE INDEX ON "Notes" ("user_id");

CREATE INDEX ON "ShareHistory" ("owner_id");

CREATE INDEX ON "ShareHistory" ("shared_id");

ALTER TABLE "ShareHistory" ADD FOREIGN KEY ("owner_id") REFERENCES "User" ("id");

ALTER TABLE "ShareHistory" ADD FOREIGN KEY ("shared_id") REFERENCES "User" ("id");

ALTER TABLE "Notes" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");