-- migrate:up
CREATE TABLE IF NOT EXISTS cryptocurrencies_rate (
    id SERIAL PRIMARY KEY,
    cryptocurrencies_rate_id INT NOT NULL REFERENCES cryptocurrencies(id),
    rate TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- migrate:down
DROP TABLE cryptocurrencies_rate;