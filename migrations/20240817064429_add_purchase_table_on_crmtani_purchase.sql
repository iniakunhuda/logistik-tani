-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_purchase.purchase (
    id INT AUTO_INCREMENT PRIMARY KEY,
    no_invoice VARCHAR(50),
    id_seller INT NOT NULL,
    id_buyer INT NULL,
    buyer_name VARCHAR(100) NULL,
    buyer_address VARCHAR(255) NULL,
    buyer_telp VARCHAR(20) NULL,
    total DECIMAL(10, 2) NOT NULL,
    purchase_date DATETIME NOT NULL,
    status ENUM('open', 'pending', 'done') NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_purchase.purchase;
-- +goose StatementEnd
