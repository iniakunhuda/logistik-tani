CREATE DATABASE IF NOT EXISTS crmtani_purchase;
USE crmtani_purchase;

CREATE TABLE IF NOT EXISTS purchase (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    no_invoice VARCHAR(100) NOT NULL,
    id_pembeli INT NOT NULL,
    id_penjual INT NOT NULL,
    total_harga INT NOT NULL,
    tanggal DATE NOT NULL,
    status ENUM('open', 'pending', 'done') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS purchase_detail (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_purchase BIGINT UNSIGNED NOT NULL,
    id_produk INT NOT NULL,
    jenis ENUM('pupuk', 'bibit', 'obat') NOT NULL,
    harga INT NOT NULL,
    qty INT NOT NULL,
    total_harga INT NOT NULL,
    tanggal DATE NOT NULL,
    is_reseller BOOLEAN NOT NULL,
    status ENUM('open', 'pending', 'done') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_purchase) REFERENCES purchase(id)
);

CREATE TABLE IF NOT EXISTS purchase_igm (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    no_invoice VARCHAR(100) NOT NULL,
    id_pembeli INT NOT NULL,
    id_penjual INT NOT NULL,
    id_produk_igm INT DEFAULT 1 NOT NULL,
    harga INT NOT NULL,
    qty INT NOT NULL,
    total_harga INT NOT NULL,
    catatan TEXT,
    pupuk_qty INT,
    pupuk_total_harga INT,
    obat_qty INT,
    obat_total_harga INT,
    bibit_qty INT,
    bibit_total_harga INT,
    status ENUM('open', 'progress', 'done') NOT NULL,
    is_deposited_to_bank BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS purchase_igm_pupuk (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_purchase BIGINT UNSIGNED NOT NULL,
    id_pupuk INT NOT NULL,
    nama_pupuk VARCHAR(255) NOT NULL,
    qty_pupuk INT NOT NULL,
    harga_pupuk INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_purchase) REFERENCES purchase_igm(id)
);

CREATE TABLE IF NOT EXISTS purchase_igm_bibit (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_purchase BIGINT UNSIGNED NOT NULL,
    id_bibit INT NOT NULL,
    nama_bibit VARCHAR(255) NOT NULL,
    qty_bibit INT NOT NULL,
    harga_bibit INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_purchase) REFERENCES purchase_igm(id)
);

CREATE TABLE IF NOT EXISTS purchase_igm_obat (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_purchase BIGINT UNSIGNED NOT NULL,
    id_obat INT NOT NULL,
    nama_obat VARCHAR(255) NOT NULL,
    qty_obat INT NOT NULL,
    harga_obat INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_purchase) REFERENCES purchase_igm(id)
);
