# Build and Run Worker

## Prerequisites: Navigate to the worker directory

```bash
cd /path/to/your/worker
```

## Commands:

### 1. Navigate to parent directory for build context

```bash
cd /your-projects  # Parent directory containing sdk-go, jobs-sdk-go, and your worker
```

### 2. Build the image using docker-compose

```bash
# From parent directory:
docker compose -f codex/automations/your-worker/docker-compose.yml build --no-cache worker
```

Or from the worker directory (using relative context):
```bash
cd codex/automations/your-worker
docker compose build --no-cache worker
```

### 3. Stop and remove old container

```bash
docker compose down worker
```

### 4. Run the updated container with enforced resource limits

```bash
docker compose --compatibility up worker -d
```

⚠️ **IMPORTANT**: The `--compatibility` flag ensures memory/CPU limits are enforced!  
Without it, the 512M memory limit in docker-compose.yml is ignored and memory leaks can affect other services.

### 5. Verify resource limits are active (CRITICAL SAFETY CHECK)

```bash
# Check if memory limit is enforced (should show 536870912 = 512MB)
docker inspect <your-worker-container-name> --format='Memory Limit: {{.HostConfig.Memory}} bytes'

# If it shows 0, limits are NOT enforced - STOP and redeploy with --compatibility flag!
```

### 6. Monitor memory usage (recommended for first 5-10 minutes after deployment)

```bash
# Real-time memory monitoring
docker stats <your-worker-container-name> --no-stream

# Or continuous monitoring:
watch -n 2 'docker stats <your-worker-container-name> --no-stream'
```

Expected behavior:
- Memory should stay under 450MB (GOMEMLIMIT triggers GC)
- If it reaches 512MB, container will be killed by OOM (protecting other services)

### 7. View logs (real-time)

```bash
docker compose logs worker -f
```

Or with limit:

```bash
docker compose logs --tail=50 worker
```

### 8. Check status

```bash
docker ps --filter "name=<your-worker-container-name>"
```

---

## Emergency Procedures

### Quick stop (if memory leak detected)

```bash
# Immediate kill
docker kill <your-worker-container-name>

# Full cleanup
docker compose down --timeout 5
```

### Check for OOM kills

```bash
# Check if container was killed due to out of memory
docker inspect <your-worker-container-name> | grep OOMKilled

# View kernel OOM logs
dmesg | grep -i "oom" | tail -20
```

---

## Safe Deployment Best Practices

### Pre-deployment checklist
- [ ] Set restart policy appropriately (`restart: "no"` for testing, `on-failure:3` for limited retries)
- [ ] Create backup image (see below)
- [ ] Ensure you have terminal access with emergency stop command ready
- [ ] Deploy during low-traffic period with team on standby
- [ ] Have monitoring dashboard open

### During deployment (first 10 minutes)
- [ ] Monitor memory usage continuously: `watch -n 2 'docker stats <your-worker-container-name> --no-stream'`
- [ ] Watch logs for errors: `docker compose logs worker -f`
- [ ] Memory should stay under 450MB
- [ ] If memory grows continuously or approaches 512MB, execute emergency stop

### Post-deployment (first hour)
- [ ] Check memory is stable
- [ ] Verify worker is processing requests
- [ ] Check no OOM kills occurred

---

## Rollback to Previous Version

### Before deploying: Create a backup

```bash
# Tag current running image as backup
docker tag worker:latest worker:backup-$(date +%Y%m%d-%H%M)
```

### To rollback:

```bash
# 1. List available backups
docker images | grep "worker"

# 2. Stop current container
docker compose down worker

# 3. Tag backup as latest
docker tag worker:backup-YYYYMMDD-HHMM worker:latest

# 4. Restart with backup image (with --compatibility for enforced limits)
docker compose --compatibility up worker -d

# 5. Verify limits are active after rollback
docker inspect <your-worker-container-name> --format='Memory Limit: {{.HostConfig.Memory}} bytes'
```

---

## ⚠️ Critical: SDK Setup

**The worker requires the Go SDK**. Without it, Docker builds will fail.

### 1. Add SDK Dependency

The SDK is available via Go modules:

```bash
go get github.com/dibbla-agents/sdk-go@latest
```

### 2. Verify go.mod

Your `go.mod` should include:

```go
require (
	github.com/dibbla-agents/sdk-go v0.0.0
	github.com/joho/godotenv v1.5.1
)
```

### 3. Update Dockerfile.worker

Edit these sections in `Dockerfile.worker`:

```dockerfile
# Update COPY paths to match your SDK location
COPY sdk-go/sdk ./sdk-go/sdk
COPY sdk-go/internal ./sdk-go/internal
COPY jobs-sdk-go ./jobs-sdk-go

# Update sed commands to match your local paths
RUN sed -i 's|/absolute/path/to/sdk-go/sdk|../sdk-go/sdk|g' go.mod && \
    sed -i 's|/absolute/path/to/sdk-go/internal|../sdk-go/internal|g' go.mod && \
    sed -i 's|/absolute/path/to/jobs-sdk-go|../jobs-sdk-go|g' go.mod
```

### 4. Build from Parent Directory

```bash
cd /your-projects  # Parent directory containing both sdk-go and worker
docker build -f codex/automations/your-worker/Dockerfile.worker -t worker:latest .
```

---

## Configuration

### Environment Variables

Required variables in `.env` file:

```env
# Worker configuration
SERVER_NAME=your-worker-name
GRPC_SERVER_ADDRESS=0.0.0.0:50051
SERVER_API_TOKEN=your-api-token

# Database configuration (if needed)
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=your-password
# DB_NAME=your-database

# OpenAI configuration (if needed)
# OPENAI_API_KEY=your-api-key

# Debug mode
CODEX_DEBUG=false
```

**Note**: The `GOMEMLIMIT=450MiB` is set in docker-compose.yml and doesn't need to be in .env

**Using defaults in docker-compose.yml**:
```yaml
environment:
  - SERVER_NAME=${SERVER_NAME:-your-worker-name}  # Uses .env value or default
  - GRPC_SERVER_ADDRESS=${GRPC_SERVER_ADDRESS:-0.0.0.0:50051}
  - CODEX_DEBUG=${CODEX_DEBUG:-false}
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

## Production Tips

**Restart Policy:**
```yaml
# Development/Testing (prevents restart loops on crashes):
restart: "no"

# Recommended: Retry once only (fail fast to prevent system overload):
restart: on-failure:1

# Limited retries (restarts up to 3 times on failure):
restart: on-failure:3

# Production (always restart unless manually stopped):
restart: unless-stopped
```

⚠️ **WARNING**: Using `restart: unless-stopped` or `restart: always` with memory leaks can create infinite crash-restart loops! 

**Best Practice**: Use `restart: on-failure:1` to retry once then fail fast. This prevents restart loops while allowing recovery from transient failures. If the container keeps failing, investigate the root cause instead of masking it with infinite restarts.

**Resource Limits:**
```yaml
deploy:
  resources:
    limits:
      cpus: '2.0'
      memory: 512M  # CRITICAL: Must use --compatibility flag to enforce
    reservations:
      cpus: '0.5'
      memory: 128M
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
- Always enforce memory limits to prevent resource exhaustion

---

## Troubleshooting

**Build fails "cannot find module":**
- Check SDK paths in `go.mod` and `Dockerfile.worker`
- Verify build context includes both SDK and worker

**Container exits immediately:**
- Check logs: `docker compose logs worker`
- Verify `.env` has required variables
- Check for OOM kills: `docker inspect <your-worker-container-name> | grep OOMKilled`

**Cannot connect to services:**
- Check environment variables
- Ensure services are on same network

**Memory limits not enforced (shows 0):**
- Must use `docker compose --compatibility up` flag
- Verify with: `docker inspect <your-worker-container-name> --format='Memory Limit: {{.HostConfig.Memory}} bytes'`

**Container killed unexpectedly:**
- Check for OOM: `docker inspect <your-worker-container-name> | grep OOMKilled`
- Review logs: `docker compose logs worker --tail=100`
- Check kernel logs: `dmesg | grep -i "oom" | tail -20`

**Container in restart loop:**
- Check restart count: `docker inspect <your-worker-container-name> --format='{{.RestartCount}}'`
- Set restart policy to "no": Change `restart: "no"` in docker-compose.yml
- Stop the loop: `docker compose down`
- Fix the underlying issue before restarting

---

## Related Guides

- [Add Worker Functions](add_worker_function.md)
- [Create Jobs](create_jobs.md)
- [Create Tasks](create_tasks.md)
- [Create Database Tables](create_database_tables.md)
