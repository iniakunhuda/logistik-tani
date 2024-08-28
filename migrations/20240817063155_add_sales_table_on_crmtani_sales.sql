-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_sales.sales (
    id INT AUTO_INCREMENT PRIMARY KEY,
    no_invoice VARCHAR(50),
    id_seller INT NOT NULL,
    id_buyer INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    sales_date DATETIME NOT NULL,
    status ENUM('open', 'pending', 'done') NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_sales.sales;
-- +goose StatementEnd
