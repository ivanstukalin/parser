-- migrate:up
CREATE TABLE IF NOT EXISTS domains (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO domains (name, url) VALUES
('bybit', 'https://api.bybit.com');

-- migrate:down
DROP TABLE domains;
