# Speech-to-Text Transcription Service

A full-stack speech-to-text transcription application that converts audio recordings into text using the Faster Whisper AI model.

## 📋 Overview

This project consists of a Go-based backend transcription service with a React frontend client, containerized using Docker for easy deployment.

### Backend Architecture

The backend is built with **Go** following clean architecture principles, featuring:

- **Transcription Service**: RESTful API built with Go that handles audio file uploads and transcription requests
- **PostgreSQL Database**: Stores transcription records and metadata
- **MinIO Object Storage**: S3-compatible storage for audio files
- **Faster Whisper AI**: Powered by [Speaches AI](https://github.com/speaches-ai/speaches) container using the **Systran/faster-whisper-small** model for a balance between quality and resource used speech recognition. The model is multilingual.
- **Dependency Injection**: Uses Wire for compile-time dependency injection
- **Clean Architecture**: Organized with separate layers for controllers, use cases, domain logic, and infrastructure

### Frontend

The frontend is a modern **React** application built with:

- **Vite** as the build tool
- **React 19** with TypeScript
- **Redux Toolkit** for state management
- **React Query** for server state management
- **Material-UI (MUI)** for UI components
- **React Router** for navigation
- **SCSS modules** for styling

Features include:
- Audio recording and file upload to get transcriptions
- Transcription history
- Protected routes with authentication

### Considerations

Currently, there's no authentication nor authorization implemented; the setup for it has been prepared, however, it shall be implemented in a future release.

Also, in the future Swagger documentation will be added to the API endpoints.

- MinIO SDK is compatible with AWS S3, so it can be used as a drop-in replacement for AWS S3.
- Fast-Whisper is compatible with OpenAI's Whisper API, so it can be used as a drop-in replacement for the Whisper API.

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.23+ (for local development)
- Node.js 20+ (for local frontend development)

### Environment Configuration

Before running the project, set up your environment files:

1. **Backend configuration:**
```bash
cd backend/transcription-service
cp .env.example .env
```

2. **Frontend configuration:**
```bash
cd frontend
cp .env.example .env
```

Edit the `.env` files as needed. The defaults should work for Docker Compose deployment.

### Running the Project

#### ⚠️⚠️ First-Time Setup (Important!) ⚠️⚠️

When running the project for the first time, you need to initialize the database before starting all services:

1. **Start the database service first:**
```bash
docker-compose up -d db
```

2. **Wait for the database to be ready** (about 10-15 seconds), then **run the migration script:**
```bash
cd backend/transcription-service
go build -o bin/migrate ./cmd/migrate
export $(cat .env | xargs) && ./bin/migrate
```

3. **Now start all services:**
```bash
cd ../..  # back to project root
docker-compose up
```

#### Subsequent Runs

After the initial setup, you can start all services normally:

```bash
docker-compose up
```

Or run in detached mode:
```bash
docker-compose up -d
```

### Accessing the Application

Once all services are running:

- **Frontend**: http://localhost:8080
- **Backend API**: http://localhost:3000
- **MinIO Console**: http://localhost:9001
- **Faster Whisper API**: http://localhost:4000

### Health Checks

All services include health check endpoints:
- Backend: http://localhost:3000/api/v1/health
- Faster Whisper: http://localhost:4000/health

## Speech Recognition Model

The application uses the **Systran/faster-whisper-small** model for transcription. This model provides:
- Fast inference time
- Good accuracy for general speech recognition
- Efficient resource usage
- Support for multiple languages

The model is automatically downloaded on the first startup of the backend service.

## 🛠️ Development

### Local Backend Development

```bash
cd backend/transcription-service
go mod download
go run cmd/main.go
```

### Local Frontend Development

```bash
cd frontend
npm install
npm run dev
```

### Running Tests

#### Backend tests
```bash
cd backend/transcription-service
go test ./...
```

#### Frontend tests
Sadly, due to time constraints, frontend tests have not been implemented yet. This will be addressed in future updates.

## Configuration

### Backend Environment Variables

Key configuration options in `backend/transcription-service/.env`:

- `APP_PORT`: Backend API port (default: 3000)
- `POSTGRES_*`: PostgreSQL connection settings
- `MINIO_*`: MinIO object storage settings
- `FASTER_WHISPER_*`: Whisper API configuration
- `FASTER_WHIPSER_MODEL`: The Whisper model to use (default: Systran/faster-whisper-small)
- `MAX_MB_FILE_SIZE`: Maximum audio file size in MB
- `ALLOWED_FILE_TYPES`: Supported audio formats

### Frontend Environment Variables

Configuration options in `frontend/.env`:

- `APP_PORT`: Frontend port (default: 8080)
- `BACKEND_URL`: Backend API URL
