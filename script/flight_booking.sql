--// Flight details
CREATE TABLE "flights" (
  "id" varchar PRIMARY KEY,	--ID
  "name" varchar(200) NOT NULL,	--Flight name
  "from" varchar(10) NOT NULL,	--From
  "to" varchar(10) NOT NULL,	--To
  "depart_date" timestamptz NOT NULL,	--Depart Time
  "status" varchar(10) NOT NULL,	--Status	(1: Active, 0: Not Active)
  "available_slot" int NOT NULL,	-- Available seats
  "created_at" timestamptz NOT NULL DEFAULT 'now()',	-- Created time
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'	-- Update time
);

--// Client Details
CREATE TABLE "clients" (
  "id" varchar PRIMARY KEY,	--ID
  "role" int,	--Role (0: Unregistered, 1: Registered, 2: Admin)
  "name" varchar(200) NOT NULL,	--Fullname
  "email" varchar(200) NOT NULL,	--Email
  "phone_number" varchar(20) NOT NULL,	--Phone
  "date_of_bith" varchar(20) NOT NULL,	-- DOB
  "identity_card" varchar(20) NOT NULL,	--Identity Number
  "address" varchar(200) NOT NULL,	--Address
  "membership_card"  varchar(20),	--Membership Number
  "password" varchar(200),	--Password
  "status" int,	--Status (0: inactive, 1: Active)
  "created_at" timestamptz NOT NULL DEFAULT 'now()',	-- Created time
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'	-- Update time
);


--// Booking Details
CREATE TABLE "bookings" (
  "id" varchar PRIMARY KEY,
  "client_id" varchar NOT NULL,	--Client Id
  "flight_id" varchar NOT NULL,	--Flight Id
  "code" varchar(20) NOT NULL,	--Booking Code
  "booked_slot" int,	-- Booked Seats
  "status" varchar(10) NOT NULL,	-- Status
  "booked_date" timestamp NOT NULL DEFAULT 'now()',	-- Booked time
  "created_at" timestamptz NOT NULL DEFAULT 'now()',	-- Created time
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'	-- Update time
);

ALTER TABLE "bookings" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");

ALTER TABLE "bookings" ADD FOREIGN KEY ("flight_id") REFERENCES "flights" ("id");
