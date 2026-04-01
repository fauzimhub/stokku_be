-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS categories CASCADE;

CREATE TABLE categories (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT
);
