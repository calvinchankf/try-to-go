CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code text NOT NULL,
    raw_data jsonb,
    created_timestamp TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_timestamp TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE UNIQUE INDEX ON orders (code);