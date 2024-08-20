-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_purchase.purchase_igm (
    id INT PRIMARY KEY,
    no_invoice VARCHAR(50),
    purchase_date DATETIME NOT NULL,
    note TEXT NULL,
    total_tebu DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    total_farmer INT NOT NULL,
    status ENUM('open', 'pending', 'done')  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_purchase.purchase_igm;
-- +goose StatementEnd
