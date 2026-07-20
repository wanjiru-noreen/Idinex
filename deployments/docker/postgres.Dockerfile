FROM postgres:16-alpine
COPY deployments/postgres/init.sql /docker-entrypoint-initdb.d/init.sql
