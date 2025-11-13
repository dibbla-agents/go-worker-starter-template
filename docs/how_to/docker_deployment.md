# Docker Deployment Guide

Deploy your worker using Docker and Docker Compose with proper SDK configuration.

## Prerequisites

- Docker 20.10+ and Docker Compose v2.0+
- Go SDK repository cloned (see below)
- `.env` file with required environment variables

---

## ⚠️ Critical: SDK Setup

**The worker requires the Go SDK**. Without it, Docker builds will fail.

### 1. Clone the SDK

```bash
git clone https://github.com/your-org/sdk-go.git
```

Recommended structure:
```
/your-projects/
  ├── sdk-go/
  └── worker_starter_template/
```

### 2. Update go.mod

Add replace directives with your local SDK paths:

```go
replace github.com/your-org/sdk-go/sdk => /absolute/path/to/sdk-go/sdk
replace github.com/your-org/sdk-go/internal => /absolute/path/to/sdk-go/internal
```

### 3. Update Dockerfile.worker

Edit these sections in `Dockerfile.worker`:

```dockerfile
# Update COPY paths to match your SDK location
COPY sdk-go/sdk ./sdk-go/sdk
COPY sdk-go/internal ./sdk-go/internal

# Update sed commands to match your local paths
RUN sed -i 's|/your/local/path/sdk-go/sdk|../sdk-go/sdk|g' go.mod && \
    sed -i 's|/your/local/path/sdk-go/internal|../sdk-go/internal|g' go.mod
```

### 4. Build from Parent Directory

```bash
cd /your-projects  # Parent directory containing both sdk-go and worker
docker build -f worker_starter_template/Dockerfile.worker -t worker:latest .
```

---

## Quick Start

### 1. Create .env file

```env
GRPC_SERVER_ADDRESS=your-grpc-server:50051
LOG_LEVEL=info
# Add other required variables
```

### 2. Build and Run

```bash
# Build
docker-compose build

# Start
docker-compose up -d

# View logs
docker-compose logs -f worker

# Stop
docker-compose down
```

---

## Configuration

### Environment Variables

Edit `.env` file or add to `docker-compose.yml`:

```yaml
environment:
  - LOG_LEVEL=debug
  - WORKER_CONCURRENCY=5
```

### Volumes

For persistent storage or file access:

```yaml
services:
  worker:
    volumes:
      - app-data:/app/data
      - ./config:/app/config:ro

volumes:
  app-data:
    driver: local
```

### Custom Networks

To connect to external services:

```yaml
networks:
  worker-network:
  external-network:
    external: true
```

---

## Common Commands

```bash
# Start
docker-compose up -d

# Stop
docker-compose down

# Restart
docker-compose restart worker

# Logs
docker-compose logs -f worker
docker-compose logs --tail=100 worker

# Status
docker-compose ps
```

---

## Production Tips

**Resource Limits:**
```yaml
deploy:
  resources:
    limits:
      cpus: '2.0'
      memory: 2G
```

**Logging:**
```yaml
logging:
  driver: "json-file"
  options:
    max-size: "10m"
    max-file: "3"
```

**Security:**
- Never commit `.env` files with secrets
- Use specific image tags (not `latest`)
- The multi-stage Dockerfile keeps images small (~10MB vs ~400MB)

---

## Example: Multi-Service Setup

Worker with nginx server sharing a volume:

```yaml
services:
  worker:
    volumes:
      - shared-files:/app/files
    networks:
      - app-network

  webserver:
    image: nginx:alpine
    ports:
      - "8081:80"
    volumes:
      - shared-files:/usr/share/nginx/html/files:ro
    depends_on:
      - worker
    networks:
      - app-network

volumes:
  shared-files:
```

---

## Troubleshooting

**Build fails "cannot find module":**
- Check SDK paths in `go.mod` and `Dockerfile.worker`
- Verify build context includes both SDK and worker

**Container exits immediately:**
- Check logs: `docker-compose logs worker`
- Verify `.env` has required variables

**Cannot connect to services:**
- Check environment variables
- Ensure services are on same network

---

## Related Guides

- [Add Worker Functions](add_worker_function.md)
- [Create Jobs](create_jobs.md)
- [Create Tasks](create_tasks.md)
- [Create Database Tables](create_database_tables.md)
