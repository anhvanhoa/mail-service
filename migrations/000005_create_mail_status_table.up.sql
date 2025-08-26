DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'type_mail_status') THEN
        CREATE TYPE type_mail_status AS ENUM ('pending', 'sent', 'delivered', 'failed', 'canceled', 'clicked', 'opened');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS mail_status (
    status type_mail_status NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO mail_status (status, name, created_by) VALUES ('pending', 'Pending', 'system');
INSERT INTO mail_status (status, name, created_by) VALUES ('sent', 'Sent', 'system');
INSERT INTO mail_status (status, name, created_by) VALUES ('delivered', 'Delivered', 'system');
INSERT INTO mail_status (status, name, created_by) VALUES ('failed', 'Failed', 'system');
INSERT INTO mail_status (status, name, created_by) VALUES ('canceled', 'Canceled', 'system');
INSERT INTO mail_status (status, name, created_by) VALUES ('clicked', 'Clicked', 'system');
INSERT INTO mail_status (status, name, created_by) VALUES ('opened', 'Opened', 'system');