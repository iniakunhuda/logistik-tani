-- +goose Up
-- +goose StatementBegin
INSERT INTO crmtani_inventory.stock_transaction (id_product_owner,id_user,stock_movement,`date`, `description`) VALUES
	 (1,1,100,'2024-08-01', 'init'),
	 (1,1,-5,'2024-08-17', 'sales'),
	 (4,3,100,'2024-08-02', 'init'),
	 (4,3,-5,'2024-08-17', 'sales'),
	 (5,4,5,'2024-08-17', 'purchase'),
	 (6,1,5,'2024-08-17', 'purchase'),
	 (3,1,80,'2024-08-02', 'init'),
	 (2,1,100,'2024-08-02', 'init'),
	 (7,5,100,'2024-08-03', 'init');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
