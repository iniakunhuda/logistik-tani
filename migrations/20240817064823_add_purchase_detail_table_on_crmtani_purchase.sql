-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_purchase.purchase_detail (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_purchase INT NOT NULL,
    id_product_owner INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    qty INT NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (id_purchase) REFERENCES crmtani_purchase.purchase(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_purchase.purchase_detail;
-- +goose StatementEnd
