#!/bin/bash
# Kubernetes Deployment Verification Script
# This script demonstrates the task-api deployment on a local Kubernetes cluster

set -e

echo "======================================"
echo "Task API Kubernetes Deployment Guide"
echo "======================================"
echo ""

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}Prerequisites:${NC}"
echo "✓ kubectl installed and configured"
echo "✓ Local Kubernetes cluster (minikube, kind, or Docker Desktop)"
echo "✓ Docker image pushed to registry"
echo ""

echo -e "${BLUE}Step 1: Create database secret${NC}"
echo "Command:"
echo "  kubectl create secret generic task-api-secrets \\"
echo "    --from-literal=database-url='postgres://user:password@db:5432/taskdb?sslmode=disable'"
echo ""

echo -e "${BLUE}Step 2: Apply Kubernetes manifests${NC}"
echo "Command:"
echo "  kubectl apply -f k8s/deployment.yaml"
echo "  kubectl apply -f k8s/service.yaml"
echo ""

echo -e "${BLUE}Step 3: Verify deployment${NC}"
echo "Commands:"
echo "  kubectl get deployments"
echo "  kubectl get pods -l app=task-api"
echo "  kubectl get svc task-api"
echo ""

echo -e "${BLUE}Step 4: Check pod status${NC}"
echo "Commands:"
echo "  kubectl rollout status deployment/task-api"
echo "  kubectl describe pod <pod-name>"
echo ""

echo -e "${BLUE}Step 5: View logs${NC}"
echo "Command:"
echo "  kubectl logs -f <pod-name>"
echo ""

echo -e "${BLUE}Step 6: Test the API${NC}"
echo "Command (local access via port-forward):"
echo "  kubectl port-forward svc/task-api 8080:80"
echo ""
echo "Then in another terminal:"
echo "  curl http://localhost:8080/health"
echo ""

echo -e "${GREEN}Deployment manifest structure:${NC}"
echo ""

echo "Deployment Details:"
echo "  - Replicas: 2"
echo "  - Image: docker.io/YOUR_DOCKER_USERNAME/task-api:latest"
echo "  - Container port: 8080"
echo "  - Database URL from secret"
echo "  - Liveness probe: /health (10s delay, 10s period)"
echo "  - Readiness probe: /health (5s delay, 5s period)"
echo "  - Resource requests: CPU 100m, Memory 128Mi"
echo "  - Resource limits: CPU 500m, Memory 512Mi"
echo ""

echo "Service Details:"
echo "  - Type: ClusterIP"
echo "  - Service port: 80"
echo "  - Target port: 8080"
echo "  - Selector: app=task-api"
echo ""

echo -e "${GREEN}CI/CD Integration:${NC}"
echo "The GitHub Actions workflow in .github/workflows/ci.yml:"
echo "  1. Runs tests and builds Docker image on every commit"
echo "  2. Pushes image to Docker Hub on main branch"
echo "  3. CD job ready for deployment trigger"
echo ""

echo "To deploy on main branch push:"
echo "  - Update the CD job in ci.yml with:"
echo "    kubectl set image deployment/task-api \\"
echo "      task-api=docker.io/YOUR_DOCKER_USERNAME/task-api:\${{ github.sha }}"
echo ""

echo -e "${GREEN}✓ Ready for deployment!${NC}"
