-- DROP DATABASE IF EXISTS crmtani_user;

CREATE DATABASE IF NOT EXISTS crmtani_inventory;
USE crmtani_inventory;

CREATE TABLE IF NOT EXISTS inventory (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    nama_produk VARCHAR(100) NOT NULL,
    hpp INT DEFAULT 0,
    harga_jual INT DEFAULT 0,
    kategori VARCHAR(100) NOT NULL,
    jenis ENUM('pupuk', 'bibit', 'obat') NOT NULL,
    stok_aktif INT DEFAULT 0,
    varietas VARCHAR(100) NOT NULL,
    status VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert dummy data into 'inventory' table
INSERT INTO inventory (id_user, nama_produk, hpp, harga_jual, kategori, jenis, stok_aktif, varietas, status)
VALUES 
(1, 'Pupuk A', 100, 150, 'Fertilizer', 'pupuk', 200, 'Variety A', 'available'),
(2, 'Bibit B', 50, 80, 'Seed', 'bibit', 100, 'Variety B', 'available'),
(1, 'Obat C', 30, 60, 'Pesticide', 'obat', 150, 'Variety C', 'available');
