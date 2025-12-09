# Documentation

Welcome to the Worker Starter Template documentation!

## Quick Start

- [Getting Started Guide](../GETTING_STARTED.md) - Quick setup and first steps

## How-To Guides

Step-by-step guides for common tasks:

- [Add a Worker Function](how_to/add_worker_function.md) - Create new external API functions
- [Create Jobs](how_to/create_jobs.md) - Build multi-step internal workflows
- [Create Tasks](how_to/create_tasks.md) - Develop reusable task operations
- [Create Database Tables](how_to/create_database_tables.md) - Set up database models and migrations
- [Add Frontend](how_to/add_frontend.md) - Add optional Vite + React web UI

## Configuration

- [Credentials Storage](credentials/README.md) - Securely store database credentials and API keys

## Planning

- [Planning Documents](planning/) - Project planning, architecture decisions, and design documents

## Deployment

- [Docker Deployment Guide](how_to/docker_deployment.md) - Production deployment with Docker

## Architecture Overview

### Components

- **Worker Functions**: Stateless functions that respond to external gRPC calls
- **Jobs**: Stateful, multi-step internal workflows with orchestration
- **Tasks**: Reusable operations that compose into jobs
- **AsyncGlobalState (ags)**: Shared resources (DB, APIs, cache, config, logger)

### When to Use What

| Component | Use When |
|-----------|----------|
| **Worker Functions** | Simple, quick operations; Webhook handlers; API integrations |
| **Jobs** | Complex orchestration; State management; Long-running processes |
| **Tasks** | Reusable operations; Need standardized error handling |

## Project Structure

```
your-project/
├── cmd/worker/              # Main application entry point
├── internal/
│   ├── worker_functions/    # External API functions
│   ├── jobs/                # Internal workflow orchestration
│   │   └── tasks/           # Reusable task operations
│   ├── models/              # Database models (GORM)
│   ├── state/               # Shared resources (AsyncGlobalState)
│   ├── config/              # Configuration management
│   └── external_api/        # External service clients
└── docs/                    # Documentation
    ├── credentials/         # Credential templates and storage
    ├── how_to/              # Step-by-step guides
    └── planning/            # Planning and design documents
```

## Support

If you have questions:
1. Check the relevant how-to guide
2. Review example implementations in the codebase
3. Open an issue for bugs or feature requests

