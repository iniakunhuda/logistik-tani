-- +goose Up
-- +goose StatementBegin
INSERT INTO crmtani_sales.sales (id,no_invoice,id_seller,id_buyer,total_price,sales_date,status) VALUES
	 (1,'SALES-001',1,4,300000.00,'2024-08-17','done'),
	 (2,'SALES-002',3,4,400000.00,'2024-08-17','done');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
