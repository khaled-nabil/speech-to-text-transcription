# speech-to-text-transcription

## Database Migrations

The transcription service uses a separate migration command for database schema management.

_Currently, the command must be run manually, but in the future it will be integrated into the build process._

### Prerequisite
Make sure you have a `.env` file in the root of the project. See `.env.example` for an example.
```bash
cp .env.example .env
```

### Running Migrations

Build the migration command:
```bash
cd backend/transcription-service
go build -o bin/migrate ./cmd/migrate
```

Run migrations up (apply all pending migrations):
```bash
export $(cat .env | xargs) && ./bin/migrate
```
