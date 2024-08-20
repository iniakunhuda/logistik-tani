-- +goose Up
-- +goose StatementBegin
INSERT INTO crmtani_user.user_land (id_user,land_name,land_address,land_area,total_obat,total_pupuk,total_bibit,total_tebu) VALUES
	 (4,'Lahan 1','Jl. Lahan 1',100.00,0.00,0.00,0.00,0.00),
	 (4,'Lahan 2','Jl. Lahan 2',150.00,0.00,0.00,0.00,0.00);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
