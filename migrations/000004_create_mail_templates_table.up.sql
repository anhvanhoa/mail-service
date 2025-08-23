DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'mail_template_status') THEN
        CREATE TYPE mail_template_status AS ENUM ('active', 'inactive');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS mail_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    subject VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    keys VARCHAR(255)[],
    provider_email VARCHAR(255) NOT NULL,
    status mail_template_status NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS idx_mail_templates_name ON mail_templates (name);
