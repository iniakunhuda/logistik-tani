-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_user.users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    role ENUM('superadmin', 'pupuk', 'bibit', 'bank', 'igm', 'petani') NOT NULL,
    telp VARCHAR(20) NOT NULL,
    email VARCHAR(255) DEFAULT NULL,
    password VARCHAR(255) DEFAULT NULL,
    saldo INT DEFAULT 0,

    token VARCHAR(255) DEFAULT NULL,
    token_expired DATETIME DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_user.users;
-- +goose StatementEnd
