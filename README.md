# Golang API POC

## MDS Entry System

A simple API for managing MDS (Master Data Service) entries.

## API Endpoints

### MDS API

- `POST /api/mds` - Create a new MDS entry
- `GET /api/mds` - Get all MDS entries
- `DELETE /api/mds/{id}` - Delete an MDS entry

## Running the Application

```bash
go run main.go
```

The server will start on port 8080.
