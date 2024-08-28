-- +goose Up
-- +goose StatementBegin
INSERT INTO crmtani_inventory.product (name,category,description,price_hpp,price_buy,price_sell) VALUES
	 ('Pupuk 1AA','pupuk','Pupuk 1AA gradenya bagus',50000,50000,60000),
	 ('Bibit Kijang Kencana','bibit','Bibit baik sekali',75000,75000,80000),
	 ('Obat Tebu Ijo','obat','Obatnya manjur',35000,35000,50000),
	 ('Bibit Diazinon 10 GR','bibit','Bibit bagus untuk tebu',60000,60000,65000),
	 ('Gula Merah','gula','Gula Merah Produksi',45000,45000,50000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
