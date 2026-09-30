# Go Worker Starter Template V0

A production-ready Go template for building worker systems with gRPC, job orchestration, and embedded frontend support.

## 🚀 Features

- **Worker Functions** - Stateless functions responding to external gRPC calls
- **Job Orchestration** - Multi-step workflows with state management
- **Task System** - Reusable, composable operations
- **Database Integration** - Optional PostgreSQL support with GORM
- **OpenAI Integration** - Ready-to-use embeddings client (optional)
- **Docker Support** - Production-ready containerization
- **Comprehensive Documentation** - Step-by-step guides for common tasks

## 🔌 Want an MCP server instead?

`_optional/mcp/` is a smaller starter for the other thing a worker can be: an
MCP server of your own, published by the platform at
`https://mcp.<your-dibbla-domain>/platform/servers/<name>` for the people your
app's access list admits. `dibbla create mcp <name>` lifts it out as a project
of its own; its [README](_optional/mcp/README.md) describes the flow and
`access_policy`. Everything else in this repository is the general worker.

## 📋 Prerequisites

- Go 1.23 or higher
- PostgreSQL (optional, for database features)
- Docker & Docker Compose (optional, for containerized deployment)

## 🎯 Quick Start

### Option 1: CLI (Recommended) 🚀
The fastest way to start. Automatically downloads the template and **renames all imports** to match your new project name.

1. **Install the tool:**
   ```bash
   go install golang.org/x/tools/cmd/gonew@latest
Create your project:

```bash
# Syntax: gonew <template-url> <your-new-module-url>
gonew github.com/dibbla-agents/go-worker-starter-template@latest github.com/your-org/your-project-name
```
Initialize Git:
The CLI creates the files but doesn't set up Git.

```bash
cd your-project-name
git init
git add .
git commit -m "Initial commit"
```
### Option 2: Standalone Project 

1. **Use this template** - Click "Use this template" button on GitHub
2. **Clone your new repository**
   ```bash
   git clone https://github.com/YOUR_USERNAME/YOUR_PROJECT_NAME.git
   cd YOUR_PROJECT_NAME
   ```
3. **Follow the setup guide** - See [GETTING_STARTED.md](GETTING_STARTED.md)

### Option 3: Add to Existing Repo (Monorepo)

To add this template inside an existing repository (e.g., `codex/automations/my-worker/`):

```bash
# 1. Navigate to your repo's subdirectory
cd codex/automations/

# 2. Clone this template with your project name
git clone https://github.com/dibbla-agents/go-worker-starter-template.git my-worker

# 3. Remove the template's git history (IMPORTANT!)
cd my-worker
rm -rf .git  # Linux/Mac
# OR: Remove-Item -Recurse -Force .git  # Windows PowerShell

# 4. Update go.mod module name to match your repo structure
# Change: module github.com/YOUR_ORG/codex/automations/my-worker

# 5. Commit to your main repo
cd ../..  # Back to codex root
git add automations/my-worker
git commit -m "Add my-worker from template"
```

**Why remove `.git`?** Each repo can only have one `.git` folder. Removing the template's `.git` makes it a regular folder that your main repo can track.

## 📚 Documentation

- **[Quick Start](QUICKSTART.md)** - Get running in 5 minutes
- **[Getting Started](GETTING_STARTED.md)** - Detailed setup guide
- **[Documentation Index](docs/README.md)** - Complete guides and references
- **[How-To Guides](docs/how_to/)** - Step-by-step tutorials for common tasks

### Key Guides

- [Add a Worker Function](docs/how_to/add_worker_function.md)
- [Create Jobs](docs/how_to/create_jobs.md)
- [Create Tasks](docs/how_to/create_tasks.md)
- [Docker Deployment](docs/how_to/docker_deployment.md)
- [Create Database Tables](docs/how_to/create_database_tables.md)

## 🏗️ Project Structure

```
├── cmd/worker/              # Application entry point
├── internal/
│   ├── worker_functions/    # Your worker function implementations
│   ├── jobs/                # Multi-step workflow definitions
│   │   └── tasks/           # Reusable task definitions
│   ├── models/              # Database models (GORM)
│   ├── state/               # Global state management
│   ├── config/              # Configuration management
│   └── embeddings/          # OpenAI embeddings client (optional)
├── docs/                    # Documentation
└── docker-compose.yml       # Multi-container setup
```

## 🔧 Customization

### Option A: Interactive Setup (Recommended)

Run the setup wizard to customize your template:

```bash
go run ./cmd/setup
```

This will ask about:
- Database support (PostgreSQL/GORM)
- Frontend (Vite + React)
- OpenAI embeddings client
- Jobs/tasks system

### Option B: Manual Setup

1. **Update module name** in `go.mod` to match your project
2. **Configure environment** - Copy `env.example` to `.env`
3. **Add your worker functions** in `internal/worker_functions/`
4. **Define your models** in `internal/models/` (if using database)
5. **Create jobs and tasks** for your specific use case

See [GETTING_STARTED.md](GETTING_STARTED.md) for detailed instructions.

## 🐳 Docker Deployment

Build and run with Docker Compose:

```bash
docker-compose up --build
```

See [Docker Deployment Guide](docs/how_to/docker_deployment.md) for production setup.

## 🧪 Development

```bash
# Format code
go fmt ./...

# Run tests
go test ./...

# Check for issues
go vet ./...

# Build
go build -o worker ./cmd/worker
```

## 📝 License

This is a template repository. When you create a project from this template, you can choose your own license.

## 🤝 Contributing

This is a template repository. Contributions to improve the template are welcome!

## 💡 Use Cases

This template is ideal for:

- Background job processors
- Webhook handlers
- API integration workers
- Data processing pipelines
- Scheduled task systems
- Microservices with complex workflows

---

**Ready to start?** Check out [GETTING_STARTED.md](GETTING_STARTED.md) for your first steps!

