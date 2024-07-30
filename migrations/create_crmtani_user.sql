-- DROP DATABASE IF EXISTS crmtani_user;

CREATE DATABASE IF NOT EXISTS crmtani_user;
USE crmtani_user;

CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    alamat VARCHAR(255),
    telp VARCHAR(15),
    role VARCHAR(15) NOT NULL,

    saldo BIGINT UNSIGNED DEFAULT 0,
    last_login VARCHAR(100) NULL DEFAULT NULL,
    alamat_kebun VARCHAR(255),
    total_obat BIGINT UNSIGNED DEFAULT 0,
    total_pupuk BIGINT UNSIGNED DEFAULT 0,
    total_bibit BIGINT UNSIGNED DEFAULT 0,
    total_tebu BIGINT UNSIGNED DEFAULT 0,
    luas_lahan BIGINT UNSIGNED DEFAULT 0,

    token VARCHAR(255) NULL DEFAULT NULL,
    token_expired TIMESTAMP NULL DEFAULT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert dummy data into 'users' table
INSERT INTO users (name, username, email, password, alamat, telp, role, saldo, last_login, alamat_kebun, total_obat, total_pupuk, total_bibit, total_tebu, luas_lahan, token, token_expired)
VALUES 
('Pupuk 1', 'pupuk1', 'pupuk1test@example.com', '$2a$14$Td7veo6J6/b4lJ/oEf/m.eI1DjOXljWVw2Z84RrPmRZnmkdm2crxK', 'Alamat', '555-1234', 'pupuk', 100000, '2024-07-29 12:00:00', 'Alamat', 10, 20, 15, 50, 100, NULL, NULL),
('Pembibit 2', 'pembibit1', 'pembibit1@example.com', '$2a$14$Td7veo6J6/b4lJ/oEf/m.eI1DjOXljWVw2Z84RrPmRZnmkdm2crxK', 'Alamat', '555-5678', 'bibit', 50000, '2024-07-28 12:00:00', 'Alamat', 5, 10, 8, 25, 50, NULL, NULL);

