CREATE DATABASE beerdb
-- ************************************** "events"

CREATE TABLE "events"
(
 "id"       int NOT NULL,
 "id_users" int NOT NULL,
 "date"     timestamp NOT NULL,
 "place"    text NOT NULL,
 CONSTRAINT "FK_26" FOREIGN KEY ( "id_users" ) REFERENCES "users" ( "id" )
);

CREATE UNIQUE INDEX "PK_events" ON "events"
(
 "id"
);

CREATE INDEX "fkIdx_26" ON "events"
(
 "id_users"
);

-- ************************************** "guests"

CREATE TABLE "guests"
(
 "id"        int NOT NULL,
 "id_events" int NOT NULL,
 "id_users"  int NOT NULL,
 "confirm"   boolean NOT NULL,
 CONSTRAINT "FK_34" FOREIGN KEY ( "id_events" ) REFERENCES "events" ( "id" ),
 CONSTRAINT "FK_37" FOREIGN KEY ( "id_users" ) REFERENCES "users" ( "id" )
);

CREATE UNIQUE INDEX "PK_guests" ON "guests"
(
 "id"
);

CREATE INDEX "fkIdx_34" ON "guests"
(
 "id_events"
);

CREATE INDEX "fkIdx_37" ON "guests"
(
 "id_users"
);

-- ************************************** "users"

CREATE TABLE "users"
(
 "id"      int NOT NULL,
 "name"    varchar(50) NOT NULL,
 "surname" varchar(50) NOT NULL,
 "email"   varchar(50) NOT NULL

);

CREATE UNIQUE INDEX "PK_users" ON "users"
(
 "id"
);












