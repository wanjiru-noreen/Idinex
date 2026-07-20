# Development Log

This log records the main development decisions, implementation steps, and verification results for the Idinex project. Each contributor should add a new entry when they introduce a feature, change the architecture, adjust data models, fix an important bug or make any change what so ever.

## Entry Format

Use this structure for every new entry:

```md
## Day N - Short Title

**Date:** YYYY-MM-DD
**Author:** Name
**Branch:** branch-name

### Goal
Briefly explain what the work was meant to achieve.

### Implementation
- Describe the files or packages changed.
- Explain the important design decisions.
- Mention any constraints or assumptions.

### Verification
- List commands run, manual checks done, or tests added.
- Note anything that could not be tested.
```

Keep entries short but useful. The goal is not to write a diary; the goal is to help the next developer understand what changed, why it changed, and how to continue without guessing.

## Day 1 - Development Infrastructure Setup

**Date:** 2026-07-20
**Author:** [Bramwel](https://github.com/dev-bramwel)
**Branch:** feat/dev-infrastructure

### Goal
Set up a repeatable development environment so a new contributor can start the full platform with one command. The work focused on local container orchestration, database bootstrapping, environment configuration, backend startup readiness, and developer workflow commands.

### Implementation
- Added a full Docker Compose stack in [docker-compose.yml](../docker-compose.yml) to start PostgreSQL, the Go backend, and the static frontend together.
- Created Dockerfiles for each service in [deployments/docker/backend.Dockerfile](../deployments/docker/backend.Dockerfile), [deployments/docker/frontend.Dockerfile](../deployments/docker/frontend.Dockerfile), and [deployments/docker/postgres.Dockerfile](../deployments/docker/postgres.Dockerfile).
- Added a PostgreSQL initialization script in [deployments/postgres/init.sql](../deployments/postgres/init.sql) so the database is bootstrapped automatically with a health check table.
- Added environment defaults in [.env.example](../.env.example) and a local [.env](../.env) file so developers can run the stack without manually creating configuration.
- Implemented backend configuration loading and database URL parsing in [backend/config/config.go](../backend/config/config.go) and [backend/config/database.go](../backend/config/database.go).
- Updated the backend entrypoint in [backend/cmd/api/main.go](../backend/cmd/api/main.go) to wait for the database before starting the server, which avoids startup race conditions during container boot.
- Added project-level Make targets in [Makefile](../Makefile) and [backend/Makefile](../backend/Makefile) for setup, development, testing, and container control.
- Added CI workflow skeletons in [.github/workflow/ci.yml](../.github/workflow/ci.yml), [.github/workflow/test.yml](../.github/workflow/test.yml), and [.github/workflow/lint.yml](../.github/workflow/lint.yml) so the new infrastructure can be exercised automatically in GitHub Actions.
- Added a basic backend config test in [backend/config/config_test.go](../backend/config/config_test.go) to validate environment overrides and database URL parsing.

### Verification
- Ran `make test` from the project root and confirmed the backend test suite passed.
- Ran `make help` to verify the new developer commands were exposed correctly.
- Parsed the Compose YAML successfully to confirm the stack definition is syntactically valid.s

