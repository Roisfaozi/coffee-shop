CREATE TABLE users
(
    id         uuid NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    username   VARCHAR,
    password   VARCHAR,
    email      VARCHAR,
    role       VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);