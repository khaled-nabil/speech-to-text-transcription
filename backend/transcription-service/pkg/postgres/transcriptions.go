package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"transcription-service/domain/transcriptionentity"
)

func (r *Repository) Save(transcription *transcriptionentity.Transcription) error {
	query := `
		INSERT INTO transcriptions (id, user_id, file_name, upload_date, transcript_text, status)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.pool.Exec(
		r.ctx,
		query,
		transcription.ID,
		transcription.UserID,
		transcription.FileName,
		transcription.UploadDate,
		transcription.TranscriptText,
		transcription.Status,
	)
	if err != nil {
		return fmt.Errorf("failed to save transcription: %w", err)
	}
	return nil
}

func (r *Repository) UpdateWithTranscription(transcription *transcriptionentity.Transcription) error {
	query := `
		UPDATE transcriptions
		SET transcript_text = $1, status = $2
		WHERE id = $3
	`
	result, err := r.pool.Exec(
		r.ctx,
		query,
		transcription.TranscriptText,
		transcription.Status,
		transcription.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update transcription: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("transcription not found")
	}
	return nil
}

func (r *Repository) GetByID(id string) (*transcriptionentity.Transcription, error) {
	query := `
		SELECT id, user_id, file_name, upload_date, transcript_text, status
		FROM transcriptions
		WHERE id = $1
	`
	var t transcriptionentity.Transcription
	err := r.pool.QueryRow(r.ctx,
		query, id).Scan(
		&t.ID,
		&t.UserID,
		&t.FileName,
		&t.UploadDate,
		&t.TranscriptText,
		&t.Status,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("transcription not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get transcription: %w", err)
	}
	return &t, nil
}

func (r *Repository) GetAllByUserID(userID string) ([]*transcriptionentity.Transcription, error) {
	query := `
		SELECT id, user_id, file_name, upload_date, transcript_text, status
		FROM transcriptions
		WHERE user_id = $1
		ORDER BY upload_date DESC
	`
	rows, err := r.pool.Query(r.ctx, query, userID)
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
			&t.Status,
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
