CREATE TABLE product
(
    id          uuid NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        VARCHAR,
    price       INTEGER,
    currency    VARCHAR,
    description VARCHAR,
    image_url   VARCHAR,
    category    VARCHAR,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);