-- Enable UUID support
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create branch_addresses table
CREATE TABLE branch_addresses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    branch_id UUID UNIQUE,
    address TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),

    -- Foreign key constraint
    CONSTRAINT fk_branch_addresses_branches FOREIGN KEY (branch_id)
        REFERENCES branches(id)
        ON DELETE SET NULL
);