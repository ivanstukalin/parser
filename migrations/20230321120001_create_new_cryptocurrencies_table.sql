-- migrate:up
CREATE TABLE IF NOT EXISTS cryptocurrencies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_cryptocurrencies_name ON cryptocurrencies(name);
CREATE INDEX IF NOT EXISTS idx_cryptocurrencies_code ON cryptocurrencies(code);

INSERT INTO cryptocurrencies (name, code) VALUES
('USDT', 'USDT'),
('Bitcoin', 'BTC'),
('Ethereum', 'ETH');

-- migrate:down
DROP INDEX IF EXISTS idx_cryptocurrencies_code;
DROP INDEX IF EXISTS idx_cryptocurrencies_name;
DROP TABLE cryptocurrencies;
