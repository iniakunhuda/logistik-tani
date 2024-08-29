-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_inventory.production (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    id_user_land INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    date_month INT NOT NULL,
    date_year INT NOT NULL,
    date_start DATE,
    date_end DATE,
    total_panen_kg INT DEFAULT 0,
    status ENUM('pending', 'done') NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_inventory.production;
-- +goose StatementEnd
