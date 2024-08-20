-- +goose Up
-- +goose StatementBegin
INSERT INTO crmtani_inventory.stock_transaction (id_product_owner,id_user,stock_movement,`date`) VALUES
	 (1,1,100,'2024-08-01'),
	 (1,1,-5,'2024-08-17'),
	 (4,3,100,'2024-08-02'),
	 (4,3,-5,'2024-08-17'),
	 (5,4,5,'2024-08-17'),
	 (6,4,5,'2024-08-17'),
	 (3,1,80,'2024-08-02'),
	 (2,1,100,'2024-08-02'),
	 (7,5,100,'2024-08-03');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
