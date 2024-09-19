-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_inventory.product_owner (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    id_product INT NOT NULL,
    stock INT NOT NULL,
    price_buy INT NOT NULL DEFAULT 0,
    price_sell INT NOT NULL DEFAULT 0,
    description TEXT,
    FOREIGN KEY (id_product) REFERENCES crmtani_inventory.product(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_inventory.product_owner;
-- +goose StatementEnd
