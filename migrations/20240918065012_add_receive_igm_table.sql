-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_purchase.receive_igm (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    no_receive VARCHAR(50) NOT NULL,
    no_spta VARCHAR(50) NOT NULL,
    no_pol VARCHAR(20) NOT NULL,
    date_out DATE NOT NULL,
    weight_kotor DECIMAL(10, 2) NOT NULL,
    weight_tara DECIMAL(10, 2) NOT NULL,
    weight_bersih DECIMAL(10, 2) NOT NULL,
    weight_rafaksi DECIMAL(10, 2),
    total_kg DECIMAL(10, 2) NOT NULL,
    harga_kg DECIMAL(10, 2) NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_purchase.receive_igm;
-- +goose StatementEnd
