-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_user.user_land (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    land_name VARCHAR(255) NOT NULL,
    land_address TEXT NOT NULL,
    land_area DECIMAL(10, 2) NOT NULL,
    total_obat DECIMAL(10, 2) NOT NULL,
    total_pupuk DECIMAL(10, 2) NOT NULL,
    total_bibit DECIMAL(10, 2) NOT NULL,
    total_tebu DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (id_user) REFERENCES crmtani_user.users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_user.user_land;
-- +goose StatementEnd
