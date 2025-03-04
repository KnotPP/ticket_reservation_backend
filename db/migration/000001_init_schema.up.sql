CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "password_hash" TEXT NOT NULL,
  "role" VARCHAR(20) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "events" (
  "id" SERIAL PRIMARY KEY,
  "organizer_id" INT NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "ticket_quota" INT NOT NULL,
  "price" DECIMAL(10,2) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "reservations" (
  "id" SERIAL PRIMARY KEY,
  "customer_id" INT NOT NULL,
  "event_id" INT NOT NULL,
  "tickets_reserved" INT NOT NULL,
  "status" VARCHAR(20) DEFAULT 'reserved',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE INDEX "idx_users_role" ON "users" ("role");

CREATE INDEX "idx_events_organizer" ON "events" ("organizer_id");

CREATE INDEX "idx_reservations_event" ON "reservations" ("event_id");

CREATE INDEX "idx_reservations_customer" ON "reservations" ("customer_id");

ALTER TABLE "events" ADD FOREIGN KEY ("organizer_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "reservations" ADD FOREIGN KEY ("customer_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "reservations" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE CASCADE;
