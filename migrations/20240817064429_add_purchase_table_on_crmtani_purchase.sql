-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_purchase.purchase (
    id INT AUTO_INCREMENT PRIMARY KEY,
    no_invoice VARCHAR(50),
    id_seller INT NULL,
    seller_name VARCHAR(100) NULL,
    seller_address VARCHAR(255) NULL,
    seller_telp VARCHAR(20) NULL,
    id_buyer INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    purchase_date DATETIME NOT NULL,
    status ENUM('open', 'pending', 'done') NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_purchase.purchase;
-- +goose StatementEnd
