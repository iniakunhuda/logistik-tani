-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_inventory.production_detail (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_production INT NOT NULL,
    id_sales_detail INT NOT NULL,
    qty_use INT NOT NULL,
    note TEXT,
    date DATE NOT NULL,
    status ENUM('pending', 'approved', 'rejected') NOT NULL,
    FOREIGN KEY (id_production) REFERENCES crmtani_inventory.production(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_inventory.production_detail;
-- +goose StatementEnd
