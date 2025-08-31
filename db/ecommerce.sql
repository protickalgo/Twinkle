CREATE DATABASE IF NOT EXISTS demo;

USE demo;

CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255),
    quantity INT,
    price INT,
    description TEXT,
    style_notes TEXT,
    size_and_fit TEXT,
    material TEXT,
    specifications TEXT,
    seller_information TEXT,
    image VARCHAR(255),  -- URL or path to product image
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
