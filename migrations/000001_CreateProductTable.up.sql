CREATE TABLE "products" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    body VARCHAR(500),
    stock INTEGER,
    created_at timestamptz,
    updated_at timestamptz
);