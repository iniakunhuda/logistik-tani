-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_inventory.stock_transaction (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_product_owner INT NOT NULL,
    id_user INT NOT NULL,
    stock_movement INT NOT NULL,  -- positive for increase, negative for decrease
    date DATETIME NOT NULL,
    FOREIGN KEY (id_product_owner) REFERENCES crmtani_inventory.product_owner(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_inventory.stock_transaction;
-- +goose StatementEnd
