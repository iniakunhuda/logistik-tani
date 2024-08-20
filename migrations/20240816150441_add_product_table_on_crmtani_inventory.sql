-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_inventory.product (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category ENUM('pupuk', 'obat', 'bibit', 'gula', 'alat') NOT NULL,
    description TEXT,
    price_hpp INT NOT NULL DEFAULT 0,
    price_buy INT NOT NULL DEFAULT 0,
    price_sell INT NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_inventory.Product;
-- +goose StatementEnd
