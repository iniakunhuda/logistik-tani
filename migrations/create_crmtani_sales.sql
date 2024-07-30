CREATE DATABASE IF NOT EXISTS crmtani_sales;
USE crmtani_sales;


CREATE TABLE IF NOT EXISTS sales (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    no_invoice VARCHAR(100) NOT NULL,
    id_penjual INT NOT NULL,
    id_pembeli INT NOT NULL,
    total_harga INT NOT NULL,
    tanggal DATE NOT NULL,
    status ENUM('open', 'pending', 'done') NOT NULL,
    is_purchased_by_igm BOOLEAN NOT NULL,
    inv_purchased_igm VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sales_detail (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_sales BIGINT UNSIGNED NOT NULL,
    id_produk INT NOT NULL,
    jenis ENUM('pupuk', 'bibit', 'obat') NOT NULL,
    harga INT NOT NULL,
    qty INT NOT NULL,
    total_harga INT NOT NULL,
    tanggal DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_sales) REFERENCES sales(id)
);

CREATE TABLE IF NOT EXISTS sales_igm (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    no_invoice VARCHAR(100) NOT NULL,
    id_penjual INT NOT NULL,
    nama_pembeli VARCHAR(255) NOT NULL,
    alamat_pembeli VARCHAR(255) NOT NULL,
    telp_pembeli VARCHAR(20) NOT NULL,
    total_harga INT NOT NULL,
    catatan TEXT,
    tanggal DATE NOT NULL,
    status ENUM('open', 'progress', 'done') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sales_igm_detail (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_sales_igm BIGINT UNSIGNED NOT NULL,
    id_produk_igm INT DEFAULT 1 NOT NULL,
    harga_beli INT NOT NULL,
    harga_jual INT NOT NULL,
    qty INT NOT NULL,
    total_harga INT NOT NULL,
    catatan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_sales_igm) REFERENCES sales_igm(id)
);


-- Insert dummy data into 'sales' table
INSERT INTO sales (no_invoice, id_penjual, id_pembeli, total_harga, tanggal, status, is_purchased_by_igm, inv_purchased_igm)
VALUES 
('INV001', 1, 2, 1500, '2024-07-29', 'done', 0, NULL),
('INV002', 2, 1, 800, '2024-07-28', 'pending', 1, 'IGM001');

-- Insert dummy data into 'sales_detail' table
INSERT INTO sales_detail (id_sales, id_produk, jenis, harga, qty, total_harga, tanggal)
VALUES 
(1, 1, 'pupuk', 150, 10, 1500, '2024-07-29'),
(2, 2, 'bibit', 80, 10, 800, '2024-07-28');

-- Insert dummy data into 'sales_igm' table
INSERT INTO sales_igm (no_invoice, id_penjual, nama_pembeli, alamat_pembeli, telp_pembeli, total_harga, catatan, tanggal, status)
VALUES 
('IGM001', 2, 'Company X', '789 Industrial Rd', '555-8765', 800, 'Urgent delivery', '2024-07-28', 'progress');

-- Insert dummy data into 'sales_igm_detail' table
INSERT INTO sales_igm_detail (id_sales_igm, id_produk_igm, harga_beli, harga_jual, qty, total_harga, catatan)
VALUES 
(1, 1, 50, 80, 10, 800, 'Delivered in good condition');