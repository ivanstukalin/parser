-- migrate:up
CREATE TABLE IF NOT EXISTS cryptocurrencies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO cryptocurrencies (name, code) VALUES
('USDT', 'USDT'),
('Bitcoin', 'BTC'),
('Ethereum', 'ETH');

-- migrate:down
DROP TABLE cryptocurrencies;
