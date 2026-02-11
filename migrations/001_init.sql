CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL,
    image_url TEXT
);
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id),
    sku VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    price_original DECIMAL(12, 2) NOT NULL,
    price_sell DECIMAL(12, 2) NOT NULL,
    provider_id VARCHAR(50),
    is_active BOOLEAN DEFAULT true
);
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id_game VARCHAR(100) NOT NULL,
    zone_id_game VARCHAR(50),
    product_id INT REFERENCES products(id),
    payment_method VARCHAR(50),
    total_price DECIMAL(12, 2),
    status VARCHAR(20) DEFAULT 'PENDING',
    payment_proof TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);