-- DROP DATABASE IF EXISTS crmtani_user;

CREATE DATABASE IF NOT EXISTS crmtani_finance;
USE crmtani_finance;

CREATE TABLE IF NOT EXISTS setup_bagi_hasil (
    id INT PRIMARY KEY AUTO_INCREMENT,
    petani INT,
    bibit INT,
    pupuk INT,
    igm INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS riwayat_pembagian (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    no_payout VARCHAR(50),
    nama_bank VARCHAR(100),
    id_sender INT,
    nama_sender VARCHAR(100),
    role_sender VARCHAR(50),
    id_receiver INT,
    nama_receiver VARCHAR(100),
    role_receiver VARCHAR(50),
    amount DECIMAL(10, 2),
    status ENUM('progress', 'done'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- insert to setup_bagi_hasil
INSERT INTO setup_bagi_hasil (petani, bibit, pupuk, igm) VALUES (50, 20, 15, 15);


-- insert riwayat_pembagian
INSERT INTO riwayat_pembagian (no_payout, nama_bank, id_sender, nama_sender, role_sender, id_receiver, nama_receiver, role_receiver, amount, status)
VALUES 
('PO001', 'Bank Central Asia', 3, 'Bank Central Asia', 'Bank', 1, 'Petani 1', 'Petani', 1000000.00, 'done'),
('PO002', 'Bank Central Asia', 3, 'Bank Central Asia', 'Bank', 2, 'Bibit B', 'Bibit', 500000.00, 'progress');