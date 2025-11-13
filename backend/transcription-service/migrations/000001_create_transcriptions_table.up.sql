CREATE TABLE IF NOT EXISTS transcriptions (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    upload_date TIMESTAMP NOT NULL,
    transcript_text TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_user_id ON transcriptions (user_id);
