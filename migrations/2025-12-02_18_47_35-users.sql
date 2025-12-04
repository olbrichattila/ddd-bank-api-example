CREATE TABLE users (
    id TEXT PRIMARY KEY, -- format: usr-[A-Za-z0-9]+
    name VARCHAR(255) NOT NULL,
    line1 VARCHAR(255) NOT NULL,
    line2 VARCHAR(255),
    line3 VARCHAR(255),
    town VARCHAR(100) NOT NULL,
    county VARCHAR(100) NOT NULL,
    postcode VARCHAR(20) NOT NULL,
    phone_number VARCHAR(20) NOT NULL, -- format: +[1-9]\d{1,14}
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);