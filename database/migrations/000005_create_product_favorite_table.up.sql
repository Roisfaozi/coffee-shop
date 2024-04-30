CREATE TABLE favorite
(
    id         uuid NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    product_id uuid REFERENCES product (id),
    user_id    uuid REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);