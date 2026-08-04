FROM golang:1.25-alpine AS build
WORKDIR /src

COPY backend/go.mod backend/go.sum* ./
RUN go mod download

COPY backend ./
RUN go build -o /bin/idinex ./cmd/api

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /bin/idinex /app/idinex
EXPOSE 8080
ENV PORT=8080
CMD ["/app/idinex"]
