-- database/schema.sql
CREATE DATABASE inventory_office;

\c inventory_office;

-- Table Categories
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table Items
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    category_id INTEGER NOT NULL,
    price DECIMAL(15, 2) NOT NULL,
    purchase_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT
);

-- Index for better performance
CREATE INDEX idx_items_category_id ON items(category_id);
CREATE INDEX idx_items_purchase_date ON items(purchase_date);
CREATE INDEX idx_items_name ON items(name);

-- Sample data
INSERT INTO categories (name, description) VALUES
('Elektronik', 'Peralatan elektronik kantor'),
('Furniture', 'Mebel dan perabotan kantor'),
('Alat Tulis', 'Perlengkapan tulis menulis');

INSERT INTO items (name, category_id, price, purchase_date) VALUES
('Laptop Dell XPS 13', 1, 15000000, '2024-06-01'),
('Monitor LG 24 inch', 1, 2500000, '2024-07-15'),
('Meja Kerja', 2, 1500000, '2024-05-10'),
('Kursi Ergonomis', 2, 2000000, '2024-05-10'),
('Printer HP LaserJet', 1, 3500000, '2024-08-01');