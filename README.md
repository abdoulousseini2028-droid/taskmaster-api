# TaskMaster API

REST API for task management built with Go, PostgreSQL, and Docker.

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

