-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_sales.sales_detail (
    id INT PRIMARY KEY,
    id_sales INT NOT NULL,
    id_product_owner INT NOT NULL,
    qty INT  NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (id_sales) REFERENCES crmtani_sales.sales(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_sales.sales_detail;
-- +goose StatementEnd
