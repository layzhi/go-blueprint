# Project test-react-plain

A Go web application built with go-blueprint using chi framework.

## Getting Started

### Prerequisites

- Go 1.21+
- Docker and Docker Compose (for database)
- Node.js 18+ and npm

### Quick Start

1. Copy `.env.example` to `.env` and configure your environment variables
2. Start the database: `make docker-run`
3. Run the application: `make run`
4. Open http://localhost:8080

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `APP_ENV` | Application environment | `local` |
| `SHUTDOWN_TIMEOUT` | Graceful shutdown timeout (seconds) | `5` |
| `FRONTEND_PORT` | React frontend port | `5173` |
| `BLUEPRINT_DB_*` | Database configuration | See `.env.example` |

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/` | Hello World |
| GET | `/health` | Health check |

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
