CREATE TABLE "User" (
  "id" BIGSERIAL PRIMARY KEY,
  "uid" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Notes" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_updated" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "User" ("id");

CREATE INDEX ON "Notes" ("user_id");

ALTER TABLE "Notes" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");