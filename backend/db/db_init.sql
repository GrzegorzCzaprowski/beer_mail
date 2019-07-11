CREATE USER testuser
    WITH PASSWORD 'md599e5ea7a6f7c3269995cba3927fd0093';

--
-- Database creation
--

CREATE DATABASE testdb
    WITH OWNER testuser;

--
--

--

\connect testdb

-- ************************************** "users"

-- *************** SqlDBM: PostgreSQL ****************;
-- ***************************************************;


-- ************************************** "public"."users"

CREATE TABLE "public"."users"
(
"id"      SERIAL NOT NULL,
"name"    varchar(50) NOT NULL,
"surname" varchar(50) NOT NULL,
"email"   varchar(50) NOT NULL,
"password"   text NOT NULL,
"admin"   boolean NOT NULL

);

CREATE UNIQUE INDEX "PK_users" ON "public"."users"
(
"id"
);
CREATE UNIQUE INDEX "email_users" ON "public"."users"
(
"email"
);


ALTER TABLE "public"."users" OWNER TO testuser;

-- *************** SqlDBM: PostgreSQL ****************;
-- ***************************************************;


-- ************************************** "public"."events"

CREATE TABLE "public"."events"
(
"id"       SERIAL NOT NULL,
"name"  varchar(50) NOT NULL,
"id_users" int NOT NULL,
"date"     timestamp NOT NULL,
"place"    text NOT NULL,
CONSTRAINT "FK_26" FOREIGN KEY ( "id_users" ) REFERENCES "public"."users" ( "id" )
);

CREATE UNIQUE INDEX "PK_events" ON "public"."events"
(
"id"
);

CREATE INDEX "fkIdx_26" ON "public"."events"
(
"id_users"
);

ALTER TABLE "public"."events" OWNER TO testuser;

-- *************** SqlDBM: PostgreSQL ****************;
-- ***************************************************;


-- ************************************** "public"."guests"

CREATE TABLE "public"."guests"
(
"id"        SERIAL NOT NULL,
"id_events" int NOT NULL,
"id_users"  int NOT NULL,
"confirm"   boolean,
CONSTRAINT "FK_34" FOREIGN KEY ( "id_events" ) REFERENCES "public"."events" ( "id" ),
CONSTRAINT "FK_37" FOREIGN KEY ( "id_users" ) REFERENCES "public"."users" ( "id" )
);

CREATE UNIQUE INDEX "PK_guests" ON "public"."guests"
(
"id"
);

CREATE INDEX "fkIdx_34" ON "public"."guests"
(
"id_events"
);

CREATE INDEX "fkIdx_37" ON "public"."guests"
(
"id_users"
);

ALTER TABLE "public"."guests" OWNER TO testuser;

--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO testuser;

--