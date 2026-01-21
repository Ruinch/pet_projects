package store

import (
	"context"
	"database/sql"
	"time"

	"forgedeploy/internal/domain"
)

type PipelineRepoPostgres struct {
	db *sql.DB
}

func NewPipelineRepo(db *sql.DB) *PipelineRepoPostgres {
	return &PipelineRepoPostgres{db: db}
}

func (r *PipelineRepoPostgres) Create(ctx context.Context, p *domain.Pipeline) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO pipelines (id, repo, commit_sha, status)
		 VALUES ($1, $2, $3, $4)`,
		p.ID, p.Repo, p.CommitSHA, p.Status,
	)
	return err
}

func (r *PipelineRepoPostgres) UpdateStatus(ctx context.Context, id string, status string) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE pipelines
		 SET status = $1, updated_at = $2
		 WHERE id = $3`,
		status, time.Now(), id,
	)
	return err
}

func (r *PipelineRepoPostgres) SaveStage(
	ctx context.Context,
	pipelineID string,
	stage string,
	status string,
	logs string,
) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO pipeline_stages
		 (pipeline_id, stage, status, logs, started_at, finished_at)
		 VALUES ($1, $2, $3, $4, now(), now())`,
		pipelineID, stage, status, logs,
	)
	return err
}
