-- Enable UUID support (jika belum)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Buat enum type untuk type field
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'branch_contact_type') THEN
        CREATE TYPE branch_contact_type AS ENUM ('home', 'phone', 'whatsapp', 'telegram');
    END IF;
END $$;

-- Create branch_contacts table
CREATE TABLE branch_contacts ( 
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    branch_id UUID,                                    
    type branch_contact_type NOT NULL,
    name VARCHAR(255),
    value VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),

    -- Foreign key constraint
    CONSTRAINT fk_nranch_contacts_branches FOREIGN KEY (branch_id)
        REFERENCES branches(id)
        ON DELETE SET NULL
);