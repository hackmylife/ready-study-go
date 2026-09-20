CREATE TABLE IF NOT EXISTS accounts (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text NOT NULL CHECK (length(btrim(name)) > 0),
    balance bigint NOT NULL CHECK (balance >= 0)
);

CREATE TABLE IF NOT EXISTS transfers (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    idempotency_key text NOT NULL UNIQUE,
    from_id bigint NOT NULL REFERENCES accounts(id),
    to_id bigint NOT NULL REFERENCES accounts(id),
    amount bigint NOT NULL CHECK (amount > 0),
    CHECK (from_id <> to_id)
);
