# Go Worker Starter Template

A production-ready Go template for building worker systems with gRPC, job orchestration, and embedded frontend support.

## 🚀 Features

- **Worker Functions** - Stateless functions responding to external gRPC calls
- **Job Orchestration** - Multi-step workflows with state management
- **Task System** - Reusable, composable operations
- **Database Integration** - Optional PostgreSQL support with GORM
- **Embedded Frontend** - Built-in web UI with Vite
- **OpenAI Integration** - Ready-to-use embeddings client
- **Docker Support** - Production-ready containerization
- **Comprehensive Documentation** - Step-by-step guides for common tasks

## 📋 Prerequisites

- Go 1.23 or higher
- PostgreSQL (optional, for database features)
- Docker & Docker Compose (optional, for containerized deployment)

## 🎯 Quick Start

1. **Use this template** - Click "Use this template" button on GitHub
2. **Clone your new repository**
   ```bash
   git clone https://github.com/YOUR_USERNAME/YOUR_PROJECT_NAME.git
   cd YOUR_PROJECT_NAME
   ```
3. **Follow the setup guide** - See [GETTING_STARTED.md](GETTING_STARTED.md)

## 📚 Documentation

- **[Getting Started](GETTING_STARTED.md)** - Quick setup and first run
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
│   ├── embeddings/          # OpenAI embeddings client
│   ├── external_api/        # External API integrations
│   └── frontend/            # Embedded frontend assets
├── frontend/                # Frontend development (Vite)
├── docs/                    # Documentation
└── docker-compose.yml       # Multi-container setup
```

## 🔧 Customization

After creating your project from this template, you should:

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

