-- +goose Up
-- +goose StatementBegin
INSERT INTO crmtani_inventory.product_owner (id_user,id_product,stock,price_buy,price_sell,description) VALUES
	 (1,1,95,50000,60000,NULL),
	 (1,3,100,35000,60000,NULL),
	 (2,1,80,65000,70000,NULL),
	 (3,2,95,75000,80000,NULL),
	 (4,1,5,60000,0,NULL),
	 (4,2,5,80000,0,NULL),
	 (5,5,100,45000,50000,NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
