-- 1. Create the database (optional)
CREATE DATABASE IF NOT EXISTS myapp;
USE myapp;

-- 2. Create the category table
CREATE TABLE category (
    category_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

-- 3. Create the sub_category table
CREATE TABLE sub_category (
    sub_category_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT,
    description TEXT,
    FOREIGN KEY (category_id) REFERENCES category(category_id) ON DELETE CASCADE
);

-- 4. Create the brand table
CREATE TABLE brand (
    brand_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

-- 5. Create the product table
CREATE TABLE product (
    product_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    category_id INT,
    sub_category_id INT,
    brand_id INT,
    FOREIGN KEY (category_id) REFERENCES category(category_id) ON DELETE SET NULL,
    FOREIGN KEY (sub_category_id) REFERENCES sub_category(sub_category_id) ON DELETE SET NULL,
    FOREIGN KEY (brand_id) REFERENCES brand(brand_id) ON DELETE SET NULL
);





-- Populate the data 
-- Insert data into the category table
INSERT INTO category (name, description)
VALUES
('Electronics', 'Devices and gadgets including phones, laptops, and other electronic equipment'),
('Appliances', 'Home and kitchen appliances');

-- Insert data into the sub_category table
INSERT INTO sub_category (name, category_id, description)
VALUES
('Smartphones', 1, 'Mobile phones with advanced features and connectivity.'),
('Laptops', 1, 'Portable computers for personal and professional use.'),
('Televisions', 1, 'LED, OLED, and Smart TVs.'),
('Refrigerators', 2, 'Home appliances for food storage and preservation.'),
('Microwaves', 2, 'Kitchen appliances for heating and cooking food.');

-- Insert data into the brand table
INSERT INTO brand (name, description)
VALUES
('Apple', 'Premium electronics brand known for iPhones, iPads, and MacBooks.'),
('Samsung', 'Leading electronics brand, offering a range of smartphones, TVs, and appliances.'),
('Sony', 'Renowned brand for TVs, audio systems, and gaming consoles.'),
('LG', 'Popular brand for home appliances and electronics.'),
('Dell', 'Major computer technology brand specializing in laptops and desktops.'),
('HP', 'Renowned for laptops, desktops, and printers.'),
('Asus', 'Known for laptops, gaming PCs, and components.'),
('Xiaomi', 'Budget-friendly electronics including smartphones and smart devices.'),
('Huawei', 'Popular for smartphones, laptops, and telecom equipment.'),
('Panasonic', 'Diverse electronics brand known for appliances and audio-visual devices.');

-- Insert data into the product table
INSERT INTO product (name, description, price, stock, category_id, sub_category_id, brand_id)
VALUES
-- Apple Products
('iPhone 14 Pro', 'Apple smartphone with 6.1-inch display, A15 Bionic chip.', 999.99, 100, 1, 1, 1),
('MacBook Air M2', 'Apple laptop with M2 chip, 13.3-inch Retina display.', 1199.99, 50, 1, 2, 1),
('iPhone 13', 'Apple smartphone with 6.1-inch display, A14 Bionic chip.', 799.99, 120, 1, 1, 1),
('MacBook Pro 16"', 'Apple laptop with M2 Max chip, 16-inch Retina display.', 2499.99, 30, 1, 2, 1),
('Apple TV 4K', 'Apple 4K streaming device with Siri remote.', 179.99, 80, 1, 3, 1),

-- Samsung Products
('Samsung Galaxy S22', 'Samsung flagship smartphone with 6.2-inch AMOLED display.', 849.99, 150, 1, 1, 2),
('Samsung QLED TV 65"', 'Samsung 65-inch QLED 4K Smart TV.', 1199.99, 60, 1, 3, 2),
('Samsung Galaxy Z Fold 4', 'Foldable smartphone with a 7.6-inch display.', 1799.99, 40, 1, 1, 2),
('Samsung Galaxy Book Pro', 'Samsung laptop with AMOLED display, Intel i7 processor.', 1299.99, 70, 1, 2, 2),
('Samsung Family Hub Refrigerator', 'Smart refrigerator with touch screen.', 2999.99, 25, 2, 4, 2),

-- Sony Products
('Sony Bravia 55"', 'Sony 55-inch OLED 4K Smart TV.', 1499.99, 45, 1, 3, 3),
('Sony WH-1000XM5', 'Noise-cancelling wireless headphones.', 349.99, 200, 1, 1, 3),
('Sony PlayStation 5', 'Sony gaming console with 4K gaming support.', 499.99, 75, 1, 3, 3),
('Sony Xperia 1 IV', 'Sony flagship smartphone with 4K display.', 999.99, 60, 1, 1, 3),
('Sony Alpha A7 IV', 'Mirrorless digital camera with 33 MP sensor.', 2499.99, 20, 1, 1, 3),

-- LG Products
('LG OLED CX 65"', 'LG 65-inch OLED 4K Smart TV with AI ThinQ.', 1899.99, 40, 1, 3, 4),
('LG Gram 17"', 'LG ultra-lightweight laptop with 17-inch display.', 1699.99, 30, 1, 2, 4),
('LG Side-by-Side Refrigerator', 'LG refrigerator with InstaView door.', 2499.99, 35, 2, 4, 4),
('LG V60 ThinQ', 'LG smartphone with 5G support and dual screen.', 799.99, 80, 1, 1, 4),
('LG NeoChef Microwave', 'Smart inverter microwave oven.', 399.99, 90, 2, 5, 4),

-- Dell Products
('Dell XPS 13', 'Dell ultrabook with Intel i7 and 13.3-inch InfinityEdge display.', 1099.99, 60, 1, 2, 5),
('Dell Alienware M15', 'Gaming laptop with NVIDIA RTX 3070.', 1799.99, 40, 1, 2, 5),
('Dell Inspiron 15', 'Dell mid-range laptop with Intel i5.', 699.99, 100, 1, 2, 5),
('Dell G5 Gaming Desktop', 'Gaming desktop with Intel i9 and RTX 3080.', 2199.99, 30, 1, 2, 5),
('Dell UltraSharp 27"', 'Dell 27-inch 4K monitor.', 749.99, 70, 1, 2, 5),

-- HP Products
('HP Spectre x360', 'Convertible laptop with 4K display.', 1399.99, 50, 1, 2, 6),
('HP Envy 13', 'Lightweight laptop with Intel i7 and 13-inch display.', 999.99, 90, 1, 2, 6),
('HP Omen 15', 'Gaming laptop with AMD Ryzen 7 and RTX 3060.', 1399.99, 50, 1, 2, 6),
('HP Pavilion 15', 'Affordable laptop with Intel i5.', 599.99, 110, 1, 2, 6),
('HP LaserJet Pro', 'Monochrome laser printer for office use.', 199.99, 80, 1, 2, 6),

-- Asus Products
('Asus ROG Zephyrus G14', 'Compact gaming laptop with AMD Ryzen 9.', 1499.99, 55, 1, 2, 7),
('Asus ZenBook 14', 'Ultra-thin laptop with OLED display.', 1199.99, 70, 1, 2, 7),
('Asus TUF Gaming A15', 'Durable gaming laptop with NVIDIA GTX 1660 Ti.', 999.99, 85, 1, 2, 7),
('Asus VivoBook S15', 'Stylish laptop with 15.6-inch display.', 799.99, 90, 1, 2, 7),
('Asus ROG Swift 27"', 'Gaming monitor with 144Hz refresh rate.', 499.99, 65, 1, 2, 7),

-- Xiaomi Products
('Xiaomi Mi 11 Ultra', 'Flagship smartphone with 108 MP camera.', 799.99, 120, 1, 1, 8),
('Xiaomi Redmi Note 10 Pro', 'Budget smartphone with 120Hz AMOLED display.', 299.99, 150, 1, 1, 8),
('Xiaomi Mi TV 4S 55"', 'Affordable 55-inch 4K Smart TV.', 499.99, 80, 1, 3, 8),
('Xiaomi Mi Band 6', 'Fitness tracker with SpO2 sensor.', 49.99, 200, 1, 1, 8),
('Xiaomi Mi Air Purifier 3H', 'Smart air purifier with OLED display.', 
