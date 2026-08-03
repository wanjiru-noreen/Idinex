# Database Schema

## Overview

The application uses PostgreSQL as its primary database. Database schema changes are managed using `golang-migrate`. Every schema change is versioned through SQL migration files located in `backend/migrations`.

## Users Table

The `users` table stores authentication and account information for registered users.

| Column | Type | Constraints |
|---------|------|-------------|
| id | UUID | Primary Key, generated with `uuid_generate_v4()` |
| username | VARCHAR(50) | NOT NULL, UNIQUE, minimum length 3 |
| email | VARCHAR(255) | NOT NULL, UNIQUE |
| password_hash | TEXT | NOT NULL |
| created_at | TIMESTAMPTZ | NOT NULL, defaults to `CURRENT_TIMESTAMP` |
| updated_at | TIMESTAMPTZ | NOT NULL, automatically updated on row modification |

## Constraints

- Primary key on `id`
- Unique constraint on `username`
- Unique constraint on `email`
- Username must contain at least three characters
- Email cannot be empty
- Password hash cannot be empty

## Indexes

PostgreSQL automatically creates unique indexes for:

- `username`
- `email`

These indexes support efficient user lookups and enforce uniqueness.

## Automatic Timestamp Updates

A PostgreSQL trigger automatically updates the `updated_at` column whenever a user record is modified.

## Migrations

Migration files are located in:

```
backend/migrations/
```

Current migrations:

- `000001_create_users.up.sql`
- `000001_create_users.down.sql`

Apply migrations:

```bash
make migrate
```

Rollback the latest migration:

```bash
make rollback
```