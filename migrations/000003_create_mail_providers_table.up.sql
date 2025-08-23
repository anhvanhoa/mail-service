DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'mail_provider_status') THEN
        CREATE TYPE mail_provider_status AS ENUM ('active', 'inactive');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS mail_providers (
    email VARCHAR(255) NOT NULL PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    port INT NOT NULL,
    host VARCHAR(255) NOT NULL,
    encryption VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type_id VARCHAR(255) NOT NULL,
    status mail_provider_status NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
)
