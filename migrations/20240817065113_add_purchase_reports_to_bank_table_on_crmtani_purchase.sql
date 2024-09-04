-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_purchase.purchase_reports_to_bank (
    id INT AUTO_INCREMENT PRIMARY KEY,
    no_report VARCHAR(255) NOT NULL,
    date_start DATE,
    date_end DATE,
    note TEXT,
    status ENUM('open', 'pending', 'done')  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_purchase.purchase_reports_to_bank;
-- +goose StatementEnd
