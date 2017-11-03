-- This file creates table for the postgreSQL

--dropdb postgres
--createdb postgres

\c postgres

CREATE TABLE userinfo
(
	uid serial                          NOT NULL,
	username    character varying(20)   NOT NULL,
	phone       character varying(50)   NOT NULL,
	homebranch  character varying(20)           ,
	CONSTRAINT  userinfo_pkey           PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);
