-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_sales.sales_igm_sugar (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_product_owner INT NOT NULL,
    id_seller INT NOT NULL,
    no_invoice VARCHAR(50),
    buyer_name VARCHAR(100) NOT NULL,
    buyer_address VARCHAR(255) NOT NULL,
    buyer_telp VARCHAR(20),
    price DECIMAL(10, 2) NOT NULL,
    qty INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    sales_date DATETIME NOT NULL,
    note TEXT NULL,
    status ENUM('open', 'pending', 'done') NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_sales.sales_igm_sugar;
-- +goose StatementEnd
