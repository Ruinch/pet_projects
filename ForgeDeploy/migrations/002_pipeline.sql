CREATE TABLE IF NOT EXISTS stages (
    id SERIAL PRIMARY KEY,
    pipeline_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    status TEXT NOT NULL,
    logs TEXT,
    started_at TIMESTAMP,
    finished_at TIMESTAMP,
    CONSTRAINT fk_pipeline
        FOREIGN KEY (pipeline_id)
        REFERENCES pipelines(id)
        ON DELETE CASCADE
);
