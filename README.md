# TaskMaster API

Production-grade REST API for team task management built with Go, PostgreSQL, and Docker.

**Performance:** Handles 200+ requests/second | **Deployment:** Single-command Docker setup

## Overview

TaskMaster is a fully containerized, horizontally-scalable task management API designed for teams. Built with Go's Gin framework and PostgreSQL, it demonstrates production-level architecture patterns including database migrations, connection pooling, and comprehensive CRUD operations.

<img width="731" height="367" alt="Screenshot 2026-01-11 at 21 56 44" src="https://github.com/user-attachments/assets/dbe61acd-0d0e-4f54-a40e-5e1624b6d03e" />


## Quick Start
```bash
# 1. Clone
git clone https://github.com/abdoulousseini2028-droid/taskmaster-api.git
cd taskmaster-api

# 2. Start database
docker-compose up -d

# 3. Run migrations
docker exec -i taskmaster-db psql -U taskmaster -d taskmaster_db < migrations/001_initial_schema.sql

# 4. Start server
go run cmd/api/main.go
```

Server runs on `http://localhost:8080`

## API Examples
```bash
# Health check
curl http://localhost:8080/health

# Create task
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"My Task","status":"todo","priority":"high"}'

# List tasks
curl http://localhost:8080/api/v1/tasks

# Get task by ID
curl http://localhost:8080/api/v1/tasks/1

# Update task
curl -X PUT http://localhost:8080/api/v1/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"done"}'

# Delete task
curl -X DELETE http://localhost:8080/api/v1/tasks/1
```

## Endpoints

- `GET /health` - Health check
- `POST /api/v1/tasks` - Create task
- `GET /api/v1/tasks` - List tasks (supports `?status=`, `?limit=`, `?offset=`)
- `GET /api/v1/tasks/:id` - Get task
- `PUT /api/v1/tasks/:id` - Update task
- `DELETE /api/v1/tasks/:id` - Delete task

## Tech Stack

- Go 1.23
- PostgreSQL 15
- Docker
- Gin framework

## Requirements

- Go 1.21+
- Docker & Docker Compose

## Architecture

```
├── cmd/api/              # Application entry point
├── internal/             # Core business logic
│   ├── handlers/         # HTTP request handlers
│   ├── models/           # Data models
│   └── database/         # Database layer
├── migrations/           # SQL migrations
└── docker-compose.yml    # Local development setup
```

The API follows clean architecture principles with clear separation between HTTP handlers, business logic, and data persistence layers.

## Features

- Full CRUD operations for task management
- Status filtering and pagination
- Priority-based task organization
- Docker containerization for consistent deployment
- PostgreSQL for reliable data persistence
- Health check endpoint for monitoring

## Development

```bash
# Run tests
go test ./...

# Run with live reload (requires air)
air
```

## Author

Built as a production-ready project to demonstrate backend systems design in Go.

