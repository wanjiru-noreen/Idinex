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

## Day 2 - Database Migration Setup and Users Schema

**Date:** 2026-08-03  
**Author:** [Bramwel](https://github.com/dev-bramwel)  
**Branch:** feat/database_schema

### Goal
Implement the initial database migration system, design the `users` table schema, and establish a maintainable workflow for future database changes.

### Implementation
- Adopted **golang-migrate** as the project's migration tool for version-controlled database schema management.
- Created the initial migration files:
  - `backend/migrations/000001_create_users.up.sql`
  - `backend/migrations/000001_create_users.down.sql`
- Designed the `users` table using **UUID** primary keys generated with `uuid_generate_v4()`.
- Added database constraints:
  - Primary key on `id`
  - Unique constraints on `username` and `email`
  - `NOT NULL` constraints for required fields
  - Validation checks for username length, non-empty email, and password hash.
- Added a PostgreSQL comment to document the `users` table.
- Implemented automatic maintenance of the `updated_at` column using a reusable PostgreSQL trigger function and trigger.
- Configured Docker Compose to include a dedicated `golang-migrate` service for running database migrations.
- Updated the root `Makefile` to support `make migrate` and `make rollback`.
- Centralized environment configuration by using `.env` and `.env.example` as the single source of truth for project configuration.
- Simplified `docker-compose.yml` to consume configuration from `.env` instead of duplicating default values.

### Design Decisions
- Selected **golang-migrate** as the project's migration framework to support versioned schema changes and reliable rollbacks.
- Chose **UUID** primary keys over auto-incrementing integers for better scalability and reduced identifier predictability.
- Used PostgreSQL **unique constraints** on `username` and `email`, which automatically create unique indexes without requiring separate `CREATE INDEX` statements.
- Implemented automatic updates to the `updated_at` column using a PostgreSQL trigger so timestamps remain accurate regardless of how records are modified.
- Centralized configuration in `.env` and `.env.example` to reduce duplication between Docker Compose, the backend, and migration tooling.

### Verification
- Verified Docker Compose successfully starts the PostgreSQL service.
- Verified the migration container can access the mounted migration files.
- Successfully executed the initial migration:

  ```bash
  docker compose run --rm migrate \
    -path=/migrations \
    -database "postgres://postgres:postgres@postgres:5432/idinex?sslmode=disable" \
    up
  ```

- Verified the `users` table migration executes successfully.
- Verified the migration tooling is integrated with Docker Compose.
- Pending:
  - Verify `make migrate`
  - Verify `make rollback`
  - Complete `docs/database.md`

