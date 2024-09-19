-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_purchase.purchase_igm_detail (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_purchase_igm INT NOT NULL,
    id_production INT NULL,
    id_user INT NOT NULL,
    id_user_land INT NULL,
    id_receive INT NULL,
    total_kg DECIMAL(10, 2) NOT NULL,
    harga_kg DECIMAL(10, 2) NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (id_purchase_igm) REFERENCES crmtani_purchase.purchase_igm(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_purchase.purchase_igm_detail;
-- +goose StatementEnd
