-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_inventory.production_detail (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_production INT NOT NULL,
    id_product_owner INT NOT NULL,
    qty_use INT NOT NULL,
    note TEXT,
    date DATETIME NOT NULL,
    FOREIGN KEY (id_production) REFERENCES crmtani_inventory.production(id),
    FOREIGN KEY (id_product_owner) REFERENCES crmtani_inventory.product_owner(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_inventory.production_detail;
-- +goose StatementEnd
