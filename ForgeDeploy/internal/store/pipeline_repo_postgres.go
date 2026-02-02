package store

import (
	"database/sql"

	"forgedeploy/internal/domain"
)

type PipelineRepoPostgres struct {
	db *sql.DB
}

func NewPipelineRepoPostgres(db *sql.DB) *PipelineRepoPostgres {
	return &PipelineRepoPostgres{db: db}
}

/* ================= PIPELINES ================= */

func (r *PipelineRepoPostgres) CreatePipeline(p *domain.Pipeline) error {
	return r.db.QueryRow(`
		INSERT INTO pipelines (name, commit_sha, status)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`,
		p.Name,
		p.CommitSHA,
		p.Status,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func (r *PipelineRepoPostgres) GetByID(id int64) (*domain.Pipeline, error) {
	var p domain.Pipeline
	err := r.db.QueryRow(`
		SELECT id, name, commit_sha, status, created_at, updated_at
		FROM pipelines WHERE id=$1
	`, id).Scan(
		&p.ID, &p.Name, &p.CommitSHA, &p.Status, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PipelineRepoPostgres) UpdatePipelineStatus(id int64, status domain.PipelineStatus) error {
	_, err := r.db.Exec(`
		UPDATE pipelines
		SET status = $1, updated_at = now()
		WHERE id = $2
	`, status, id)
	return err
}

func (r *PipelineRepoPostgres) GetAll() ([]*domain.Pipeline, error) {
	rows, err := r.db.Query(`
		SELECT id, name, commit_sha, status, created_at, updated_at
		FROM pipelines
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pipelines []*domain.Pipeline
	for rows.Next() {
		var p domain.Pipeline
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.CommitSHA,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		pipelines = append(pipelines, &p)
	}

	return pipelines, nil
}

/* ================= STAGES ================= */

func (r *PipelineRepoPostgres) CreateStage(s *domain.Stage) error {
	return r.db.QueryRow(`
		INSERT INTO stages (pipeline_id, name, status)
		VALUES ($1, $2, $3)
		RETURNING id
	`, s.PipelineID, s.Name, s.Status).Scan(&s.ID)
}

func (r *PipelineRepoPostgres) UpdateStage(s *domain.Stage) error {
	_, err := r.db.Exec(`
		UPDATE stages
		SET
			status = $1,
			logs = $2,
			started_at = $3,
			finished_at = $4
		WHERE id = $5
	`,
		s.Status,
		s.Logs,
		s.StartedAt,
		s.FinishedAt,
		s.ID,
	)
	return err
}

func (r *PipelineRepoPostgres) GetStages(pipelineID int64) ([]*domain.Stage, error) {
	rows, err := r.db.Query(`
		SELECT id, pipeline_id, name, status, logs, started_at, finished_at
		FROM stages
		WHERE pipeline_id=$1
		ORDER BY id
	`, pipelineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []*domain.Stage
	for rows.Next() {
		var s domain.Stage
		if err := rows.Scan(
			&s.ID,
			&s.PipelineID,
			&s.Name,
			&s.Status,
			&s.Logs,
			&s.StartedAt,
			&s.FinishedAt,
		); err != nil {
			return nil, err
		}
		stages = append(stages, &s)
	}
	return stages, nil
}
