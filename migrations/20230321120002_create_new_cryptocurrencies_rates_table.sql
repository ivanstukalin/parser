-- migrate:up
CREATE TABLE IF NOT EXISTS cryptocurrencies_rate (
    id SERIAL PRIMARY KEY,
    cryptocurrencies_rate_id INT NOT NULL REFERENCES cryptocurrencies(id),
    rate TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_cryptocurrencies_rate_crypto_id ON cryptocurrencies_rate(cryptocurrencies_rate_id);

-- migrate:down
DROP TABLE cryptocurrencies_rate;