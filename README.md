# Netgraph2 Example

This repository contains a simple Go server using `gorilla/mux` and a Svelte frontend.

## Requirements
- Go 1.20+
- Node.js 18+

## Building the frontend
```
cd frontend
npm install
npm run build
```
This creates the production files in `frontend/dist`.

## Running the backend
```
go run ./cmd/server
```
The server listens on `http://localhost:8080` and serves the frontend files.

## API Example
Visit `http://localhost:8080/api/hello` to see a sample JSON response.
