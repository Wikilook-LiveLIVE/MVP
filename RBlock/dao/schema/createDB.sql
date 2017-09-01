CREATE DATABASE rblock
WITH
OWNER = postgres
ENCODING = 'UTF8'
LC_COLLATE = 'English_United States.1252'
LC_CTYPE = 'English_United States.1252'
TABLESPACE = pg_default
CONNECTION LIMIT = -1;

COMMENT ON DATABASE rblock
IS 'R_Block project database';

CREATE SCHEMA IF NOT EXISTS rblock;


