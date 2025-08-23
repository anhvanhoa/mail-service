DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'mail_history_status') THEN
        CREATE TYPE mail_history_status AS ENUM ('pending', 'sent', 'delivered', 'failed', 'canceled', 'clicked', 'opened');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS mail_histories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    template_id VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    tos VARCHAR(255)[],
    data JSONB NOT NULL,
    email_provider VARCHAR(255) NOT NULL,
    status mail_history_status NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS idx_mail_histories_mail_id ON mail_histories (mail_id);
CREATE INDEX IF NOT EXISTS idx_mail_histories_status ON mail_histories (status);