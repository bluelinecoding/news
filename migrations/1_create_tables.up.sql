CREATE TABLE IF NOT EXISTS feeds (
    id text PRIMARY KEY,
    category text NOT NULL,
    provider text NOT NULL,
    url text NOT NULL,
    created_at timestamp DEFAULT NOW()
)
