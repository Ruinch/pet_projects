CREATE TABLE pipelines (
    id TEXT PRIMARY KEY,
    repo TEXT NOT NULL,
    commit_sha TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ
);

CREATE TABLE pipeline_stages (
    id SERIAL PRIMARY KEY,
    pipeline_id TEXT NOT NULL REFERENCES pipelines(id) ON DELETE CASCADE,
    stage TEXT NOT NULL,
    status TEXT NOT NULL,
    logs TEXT,
    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ
);
