CREATE TABLE IF NOT EXISTS pipelines (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    commit_sha TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);
