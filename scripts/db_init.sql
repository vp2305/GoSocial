-- Connect to the default database
\c postgres

-- Create the database if it does not exist
DO
$$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'social') THEN
      PERFORM dblink_exec('dbname=postgres', 'CREATE DATABASE social');
   END IF;
END
$$;

-- Connect to the new database
\c social

-- Create the extension if it does not exist
CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS pg_trgm;