-- Enable UUID support
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create categories table
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

-- Index for soft deletes
CREATE INDEX idx_categories_deleted_at ON categories(deleted_at);