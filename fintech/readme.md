# Fintech - Financial Technology Platform

## Overview

A financial application built with modern Go practices, leveraging Temporal for workflow orchestration, TigerBeetle for financial transactions, and PostgreSQL for data persistence.

## Tech Stack

- **Temporal**: Workflow engine for orchestrating complex financial processes
- **TigerBeetle**: High-performance financial accounting database
- **PostgreSQL**: Relational database for user and transaction data
- **Fiber**: Web framework for HTTP APIs

## Project Structure

```
fintech/
├── cmd/              # Main server entry point
├── internal/
│   ├── config/       # Configuration settings
│   ├── handlers/     # HTTP request handlers
│   │   ├── auth.go   # Authentication endpoints
│   │   ├── transaction.go # Transaction endpoints
│   │   └── user.go   # User management endpoints
│   ├── helpers/      # Utility functions (body parsing, error handling)
│   └── models/       # Data models and database logic
└── docker-compose.yml # Local development services
```

## Features

- User registration and authentication
- Financial transaction processing
- Workflow orchestration with Temporal
- Secure financial data storage with TigerBeetle

## Getting Started

### Prerequisites

- Go 1.21+
- Docker and Docker Compose
- Temporal server (or use Docker Compose)
- TigerBeetle

### Running Locally

1. Start the required services:
   ```bash
   docker-compose up -d
   ```

2. Run the application:
   ```bash
   go run cmd/main.go
   ```
