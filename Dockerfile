# Build frontend
FROM node:18 AS frontend-build
WORKDIR /frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci
COPY frontend/ .
RUN npm run build

# Build Go binary
FROM golang:1.22 AS backend-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
RUN CGO_ENABLED=0 go build -o server ./cmd/server

# Final runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=backend-build /app/server ./server
COPY --from=frontend-build /frontend/dist ./frontend/dist
EXPOSE 8080
CMD ["./server"]
