CREATE TABLE transactions (
    id TEXT PRIMARY KEY, -- format: tan-[A-Za-z0-9]
    account_number CHAR(8) NOT NULL REFERENCES accounts(account_number) ON DELETE CASCADE,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount NUMERIC(12, 2) NOT NULL CHECK (amount >= 0.00),
    currency CHAR(3) NOT NULL DEFAULT 'GBP',
    type VARCHAR(50) NOT NULL CHECK (type IN ('deposit', 'withdrawal')),
    reference TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transactions_account_number ON transactions(account_number);
CREATE INDEX idx_transactions_user_id ON transactions(user_id);