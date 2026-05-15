# CI/CD Pipeline Setup Guide

## GitHub Actions Workflow

The CI/CD pipeline has been created at `.github/workflows/ci.yml` with two jobs:

### CI Job
- **Triggers**: On push to `main` or any pull request
- **Steps**:
  1. Checks out code
  2. Sets up Go 1.21 environment
  3. Runs `go vet` for static analysis
  4. Runs `go test -race` for race condition detection
  5. Builds Docker image with git SHA and latest tags
  6. (On main push only) Logs into Docker Hub
  7. (On main push only) Pushes images to Docker Hub

### CD Job
- **Triggers**: Only on successful CI completion for main branch pushes
- **Steps**:
  1. Checks out code
  2. Placeholder for Kubernetes deployment (ready for real deployment command)

## Prerequisites

### GitHub Repository Secrets
Add these secrets to your GitHub repository (Settings → Secrets and variables → Actions):
- `DOCKER_USERNAME`: Your Docker Hub username
- `DOCKER_PASSWORD`: Your Docker Hub personal access token (not your password!)

### Dockerfile
Ensure you have a `Dockerfile` in the root of your repository. Example:

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /build/api .

EXPOSE 8080
CMD ["./api"]
```

## Kubernetes Deployment

### Prerequisites
- Kubernetes cluster running (minikube, Docker Desktop, or cloud provider)
- Docker image pushed to Docker Hub or accessible registry
- `kubectl` configured with your cluster context

### Database Secret
Create a Kubernetes secret for your database URL:

```bash
kubectl create secret generic task-api-secrets \
  --from-literal=database-url='postgres://user:password@db-host:5432/taskdb?sslmode=disable'
```

### Deploy to Kubernetes

```bash
# Apply all manifests
kubectl apply -f k8s/

# Check deployment status
kubectl rollout status deployment/task-api

# View pods
kubectl get pods -l app=task-api

# Port forward to test locally
kubectl port-forward svc/task-api 8080:80
```

### Verify Deployment

```bash
# Check service
kubectl get svc task-api

# Check deployment
kubectl get deployment task-api

# View logs from a pod
kubectl logs <pod-name>

# Describe deployment for events
kubectl describe deployment task-api
```

## Image Configuration

Update the image path in `k8s/deployment.yaml`:

Replace `YOUR_DOCKER_USERNAME` with your actual Docker Hub username:

```yaml
image: docker.io/YOUR_DOCKER_USERNAME/task-api:latest
```

Or if using a different registry (e.g., GitHub Container Registry):

```yaml
image: ghcr.io/YOUR_GITHUB_USERNAME/task-api:latest
```

## API Health Check

The deployment includes liveness and readiness probes that check `/health` endpoint. Ensure your Go API has this endpoint implemented:

```go
router.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "healthy"})
})
```

## Production Deployment Steps

When ready for production:

1. **Update CD job**: Uncomment and customize the `kubectl set image` command in the CD job to actually deploy
2. **Add ingress**: Create an Ingress resource for external access
3. **Configure TLS**: Add HTTPS/TLS certificates
4. **Environment management**: Use multiple namespaces and environment-specific overlays with Kustomize or Helm
5. **Monitoring**: Add Prometheus metrics and alerting
6. **Logging**: Integrate with ELK, Datadog, or similar logging services
