CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   username VARCHAR(50) UNIQUE NOT NULL,
   password_hash TEXT NOT NULL,
   role VARCHAR(20) NOT NULL
);

CREATE TABLE motorcycles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    brand VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    motorcycle_id INT REFERENCES motorcycles(id) ON DELETE CASCADE,
    order_date TIMESTAMP DEFAULT NOW(),
    total_price DECIMAL(10, 2) NOT NULL
);
