CREATE TABLE profile
(
    id           uuid NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name   VARCHAR,
    last_name    VARCHAR,
    display_name VARCHAR,
    gender       VARCHAR,
    address      VARCHAR,
    user_id      uuid UNIQUE REFERENCES users (id),
    phone_number INTEGER,
    birthday     DATE,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
