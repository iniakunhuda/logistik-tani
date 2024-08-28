-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_purchase.purchase_reports_to_bank (
    id INT AUTO_INCREMENT PRIMARY KEY,
    date_start DATE,
    date_end DATE,
    note TEXT,
    status VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_purchase.purchase_reports_to_bank;
-- +goose StatementEnd
