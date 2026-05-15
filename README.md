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
- **CI/CD Pipeline**: Automated testing, building, and Docker image publishing via GitHub Actions
- **Kubernetes Ready**: Complete Kubernetes deployment manifests for scalable production deployment

## CI/CD & Deployment

### GitHub Actions Pipeline

Automated CI/CD pipeline configured in `.github/workflows/ci.yml`:
- **Continuous Integration**: Runs on every push and pull request
  - Go static analysis (`go vet`)
  - Unit tests with race condition detection (`go test -race`)
  - Docker image build and push to Docker Hub
- **Continuous Deployment**: Runs on successful CI completion for main branch
  - Ready for automated deployment to Kubernetes

**Required Secrets** (GitHub Settings → Secrets):
- `DOCKER_USERNAME`: Docker Hub username
- `DOCKER_PASSWORD`: Docker Hub personal access token

### Kubernetes Deployment

Production-grade Kubernetes manifests in `k8s/`:

**Deployment** (`k8s/deployment.yaml`):
- 2 replicas for high availability
- Auto-scaling ready (liveness & readiness probes)
- Database connection via ConfigMap/Secret
- Resource limits (CPU: 100m-500m, Memory: 128Mi-512Mi)
- Health check probes at `/health` endpoint

**Service** (`k8s/service.yaml`):
- ClusterIP service exposing port 80 → 8080
- Load balancing across replicas

### Quick Kubernetes Deploy

```bash
# 1. Create database secret
kubectl create secret generic task-api-secrets \
  --from-literal=database-url='postgres://user:pass@db:5432/taskdb?sslmode=disable'

# 2. Update image in k8s/deployment.yaml
sed -i 's/YOUR_DOCKER_USERNAME/your-username/' k8s/deployment.yaml

# 3. Apply manifests
kubectl apply -f k8s/

# 4. Verify deployment
kubectl rollout status deployment/task-api
kubectl get pods -l app=task-api

# 5. Test locally
kubectl port-forward svc/task-api 8080:80
curl http://localhost:8080/health
```

See [CICD_SETUP.md](CICD_SETUP.md) for detailed setup instructions.
See [deploy.sh](deploy.sh) for deployment verification guide.

## Development

```bash
# Run tests
go test ./...

# Run with live reload (requires air)
air
```

## Author

Built as a production-ready project to demonstrate backend systems design in Go.

