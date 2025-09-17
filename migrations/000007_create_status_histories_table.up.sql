CREATE TABLE IF NOT EXISTS status_histories (
    status type_mail_status NOT NULL,
    mail_history_id VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (status, mail_history_id)
)