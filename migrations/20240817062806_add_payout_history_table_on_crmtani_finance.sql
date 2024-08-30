-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_finance.payout_history (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_sender INT NOT NULL,
    id_receiver INT NOT NULL,
    no_invoice VARCHAR(255) NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    bank_note TEXT,
    id_purchase_reports_to_bank INT,
    date_payout DATE,
    status ENUM('pending', 'approved', 'rejected') NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    rejected_date TIMESTAMP NULL,
    rejected_message TEXT NULL,
    approved_message TEXT NULL,
    approved_date TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_finance.payout_history;
-- +goose StatementEnd
