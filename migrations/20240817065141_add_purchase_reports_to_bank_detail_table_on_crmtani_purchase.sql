-- +goose Up
-- +goose StatementBegin
CREATE TABLE crmtani_purchase.purchase_reports_to_bank_detail (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_purchase_reports_to_bank INT,
    id_purchase_igm INT,
    FOREIGN KEY (id_purchase_reports_to_bank) REFERENCES crmtani_purchase.purchase_reports_to_bank(id),
    FOREIGN KEY (id_purchase_igm) REFERENCES crmtani_purchase.purchase_igm(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE crmtani_purchase.purchase_reports_to_bank_detail;
-- +goose StatementEnd
