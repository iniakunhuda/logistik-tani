-- +goose Up
-- +goose StatementBegin
INSERT INTO crmtani_user.`users` (id,name,address,`role`,telp,email,password,saldo, token, token_expired) VALUES
	 (1, 'Pupuk 1','Jl. Pupuk No. 1','pupuk','08111','pupuk1@gmail.com','$2y$12$V79nT56atGXTSCBzbbMJFurJqAe5NPrvrVwGfAyQjoGDGQb0GA3h2',0.00, null, null),
	 (2, 'Pupuk 2','Jl. Pupuk No. 2','pupuk','08112','pupuk2@gmail.com','$2y$12$MJ2wKqeCAr9eh3dQhbWXpuPiyoT.ts88Jbx.z.L3UAd0BarhrNwKa',0.00, null, null),
	 (3, 'Pembibit 1','Jl. Pogot','bibit','08221','bibit1@gmail.com','$2y$12$4A0LLjXyXEu5t4MxtWBve.aGW69XarPjOWkpeRU72ysYv5BEZCVG6',0.00, null, null),
	 (4, 'Petani 1','Jl. Habibie','petani','08331','petani1@gmail.com','$2y$12$vsbLi2AtwdJDNMf893Dl7OfclxejhDByC1X3/tqsedm47fIYK1uru',0.00, null, null),
	 (5, 'IGM','IGM','igm','08991','igm1@gmail.com','$2y$12$9zK2Zz3a/z2Iq6GX62yvMO/M5inZJVsPsFm6r6rdUwv4ri2k3XkzS',0.00, null, null),
	 (6, 'Bank','Bank','bank','08992','bank1@gmail.com','$2y$12$OZjMhn68hMefsDsCTEcTluAkduN4j2cTQ.zYkF/H6GxnUVDIxs8f2',0.00, null, null);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
