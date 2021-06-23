-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Notes:
-- * Assuming the X, Y coordinates have a max. lenght that fits into postgres type. 
--   If it wouldn't, we'd have to take a different approach.
-- * Sometimes it's a good idea to have soft-delete, we'd add this attribute if desired:
--   `deleted_at timestamp with time zone NULL DEFAULT NULL`

CREATE TABLE IF NOT EXISTS ports (
    "uuid" uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    key text NOT NULL UNIQUE,
    name text DEFAULT NULL,
    city text DEFAULT NULL,
    country text DEFAULT NULL,
    province text DEFAULT NULL,
    timezone text DEFAULT NULL,
    code text DEFAULT NULL,
    alias varchar[], 
    regions varchar[],
    unlocs varchar[],
    xcoord text DEFAULT NULL,
    ycoord text DEFAULT NULL
);

