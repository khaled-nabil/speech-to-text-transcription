package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"transcription-service/domain/persistance"
	"transcription-service/domain/transcriptionentity"
	"transcription-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(config *config.Config) (persistance.TranscriptionRepository, error) {
	ctx := context.Background()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.Database,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Repository{pool: pool}, nil
}

func (r *Repository) Save(ctx context.Context, transcription *transcriptionentity.Transcription) error {
	query := `
		INSERT INTO transcriptions (id, user_id, file_name, upload_date, transcript_text)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.pool.Exec(
		ctx,
		query,
		transcription.ID,
		transcription.UserID,
		transcription.FileName,
		transcription.UploadDate,
		transcription.TranscriptText,
	)
	if err != nil {
		return fmt.Errorf("failed to save transcription: %w", err)
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*transcriptionentity.Transcription, error) {
	query := `
		SELECT id, user_id, file_name, upload_date, transcript_text
		FROM transcriptions
		WHERE id = $1
	`
	var t transcriptionentity.Transcription
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&t.ID,
		&t.UserID,
		&t.FileName,
		&t.UploadDate,
		&t.TranscriptText,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("transcription not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get transcription: %w", err)
	}
	return &t, nil
}

func (r *Repository) GetAllByUserID(ctx context.Context, userID string) ([]*transcriptionentity.Transcription, error) {
	query := `
		SELECT id, user_id, file_name, upload_date, transcript_text
		FROM transcriptions
		WHERE user_id = $1
		ORDER BY upload_date DESC
	`
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transcriptions: %w", err)
	}
	defer rows.Close()

	var transcriptions []*transcriptionentity.Transcription
	for rows.Next() {
		var t transcriptionentity.Transcription
		if err = rows.Scan(
			&t.ID,
			&t.UserID,
			&t.FileName,
			&t.UploadDate,
			&t.TranscriptText,
		); err != nil {
			return nil, fmt.Errorf("failed to scan transcription: %w", err)
		}
		transcriptions = append(transcriptions, &t)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transcriptions: %w", err)
	}

	return transcriptions, nil
}

func (r *Repository) Close() error {
	if r.pool != nil {
		r.pool.Close()
	}
	return nil
}
