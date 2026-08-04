DOCKER_COMPOSE := $(shell if command -v docker-compose >/dev/null 2>&1; then echo docker-compose; elif docker compose version >/dev/null 2>&1; then echo "docker compose"; else echo "docker compose"; fi)

MIGRATE_IMAGE := migrate/migrate:v4.18.3

DOCKER_NETWORK := $(shell $(DOCKER_COMPOSE) ps -q postgres | xargs docker inspect -f '{{range $$k,$$v := .NetworkSettings.Networks}}{{$$k}}{{end}}')

ifneq (,$(wildcard .env))
	include .env
	export
endif

.PHONY: help setup dev stop restart build test lint fmt migrate rollback seed reset-db clean logs backend frontend db docker-build docker-push deploy backup restore

help:
	@echo "Available commands:"
	@printf '%-15s %s\n' 'setup' 'First-time project setup'
	@printf '%-15s %s\n' 'dev' 'Start the full development stack'
	@printf '%-15s %s\n' 'test' 'Run the backend test suite'
	@printf '%-15s %s\n' 'stop' 'Stop the development environment'
	@printf '%-15s %s\n' 'restart' 'Restart the development environment'
	@printf '%-15s %s\n' 'logs' 'Follow application logs'
	@printf '%-15s %s\n' 'backend' 'Start the backend service only'
	@printf '%-15s %s\n' 'frontend' 'Start the frontend service only'
	@printf '%-15s %s\n' 'db' 'Start PostgreSQL only'

setup:
	@if [ ! -f .env ]; then cp .env.example .env; fi
	@$(DOCKER_COMPOSE) build
	@$(DOCKER_COMPOSE) up -d postgres
	@echo "Development environment is ready. Run 'make dev' to start the stack."

dev:
	$(DOCKER_COMPOSE) up --build

stop:
	$(DOCKER_COMPOSE) down

restart: stop dev

build: docker-build

test:
	cd backend && go test ./...

lint:
	cd backend && golangci-lint run

fmt:
	cd backend && go fmt ./...

format-check:
	cd backend && find . -name "*.go" -print0 | xargs -0 gofmt -w
	cd backend && git diff --exit-code

migrate: db
	$(DOCKER_COMPOSE) run --rm migrate \
		-path=/migrations \
		-database "$(DATABASE_URL)" \
		up

rollback: db
	$(DOCKER_COMPOSE) run --rm migrate \
		-path=/migrations \
		-database "$(DATABASE_URL)" \
		down 1

seed:
	@echo "The PostgreSQL container bootstraps the development database automatically."

reset-db:
	$(DOCKER_COMPOSE) down -v
	$(DOCKER_COMPOSE) up -d postgres

clean:
	$(DOCKER_COMPOSE) down -v
	docker system prune -f

logs:
	$(DOCKER_COMPOSE) logs -f

backend:
	$(DOCKER_COMPOSE) up --build backend

frontend:
	$(DOCKER_COMPOSE) up --build frontend

db:
	$(DOCKER_COMPOSE) up -d postgres

docker-build:
	$(DOCKER_COMPOSE) build

docker-push:
	@echo "Docker push is not configured for local development."

deploy:
	@echo "Deployment targets are not configured yet."

backup:
	$(DOCKER_COMPOSE) exec -T postgres pg_dump -U $${POSTGRES_USER:-postgres} $${POSTGRES_DB:-idinex} > backup.sql

restore:
	$(DOCKER_COMPOSE) exec -T postgres psql -U $${POSTGRES_USER:-postgres} -d $${POSTGRES_DB:-idinex} < backup.sql
