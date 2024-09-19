-- +goose Up
-- +goose StatementBegin
INSERT INTO crmtani_sales.sales_detail (id,id_sales,id_product_owner,qty,price,subtotal) VALUES
	 (1,1,1,5,60000.00,300000.00),
	 (2,2,4,5,80000.00,400000.00);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
