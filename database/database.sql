CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
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

CREATE TABLE size
(
    id        uuid NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    size_name VARCHAR
);
CREATE TABLE product_size
(
    product_id uuid REFERENCES product (id),
    size_id    uuid REFERENCES size (id),
    PRIMARY KEY (product_id, size_id)
);

CREATE TABLE favorite
(
    id         uuid NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    product_id uuid REFERENCES product (id),
    user_id    uuid REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (username, email, role) VALUES ('coba1', 'coba1@example.com', 'coba1');

truncate TABLE product cascade;
truncate TABLE size cascade;
truncate TABLE product_size cascade;
truncate TABLE favorite cascade;
truncate TABLE users cascade;

select * from product where id = '4fbbc3a2-8516-4d9b-a5e5-5b64790237d6';
select * from users;

insert into size(size_name) values ('R');
insert into size(size_name) values ('L');
insert into size(size_name) values ('XL');

SELECT p.id, p.name, p.price, p.currency, p.description, p.image_url, p.created_at, p.updated_at,
       s.id as size_id, s.size_name
FROM product p
         LEFT JOIN product_size ps ON p.id = ps.product_id
         LEFT JOIN size s ON ps.size_id = s.id
WHERE p.id = 'f573f841-c04b-4af3-9a46-e79f4ac94d68';

SELECT p.id, p.name, p.price, p.currency, p.description, p.image_url, p.category, p.created_at, p.updated_at,
       s.id as size_id, s.size_name
FROM product p
         LEFT JOIN product_size ps ON p.id = ps.product_id
         LEFT JOIN size s ON ps.size_id = s.id;
SELECT DISTINCT p.id, p.name, p.price, p.currency, p.description, p.image_url, p.category, p.created_at, p.updated_at,
                s.id as size_id, s.size_name
FROM product p
         LEFT JOIN product_size ps ON p.id = ps.product_id
         LEFT JOIN size s ON ps.size_id = s.id
WHERE p.category ILIKE '%makanan%'
ORDER BY p.id LIMIT 10 OFFSET 2;

SELECT id, username, "role", "password" FROM users WHERE username = 'Testing-1713977486';