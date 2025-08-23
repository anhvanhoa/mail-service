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
    updated_at TIMESTAMP NULL
);