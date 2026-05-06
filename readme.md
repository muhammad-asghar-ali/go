# Project Overview

This repository contains several independent projects, each with its own functionality. Below is an overview of each project and its structure.

## Projects

### [ai-fun](./ai-fun)

This project is designed to explore and implement various artificial intelligence models and algorithms.

### [authis](./authis)

A JWT-based authentication system built with Fiber, Redis, and PostgreSQL.

### [bitora](./bitora)

A blockchain implementation with proof-of-work, using BoltDB for storage and a CLI interface.

### [bookstore](./bookstore)

A simple bookstore application with MySQL database and GORM.

- **cmd**: Main GO server.
- **internal**:
  - **config**: Configuration settings.
  - **handlers**: HTTP request handlers.
  - **models**: Database models.
  - **routes**: API route definitions.
  - **utils**: Utility functions and helpers.

### [bootcamp](./bootcamp)

Go learning exercises covering fundamental concepts.

- **01_hello**: Hello world program.
- **02_variables**: Variables and data types.
- **03_user_input**: Reading user input.
- **04_conversions**: Type conversions.
- **05_times**: Working with time.
- **06_pointers**: Pointers in Go.
- **07_arrays**: Arrays.
- **08_slices**: Slices.

### [cache](./cache)

A caching service implementing an LRU cache using a doubly-linked list and hash map.

- **internal**: Core logic for the caching mechanism, including queue operations.

### [chronos](./chronos)

Chronos is a bot application that integrates with Slack and provides various functionalities like age calculation, system monitoring, and more. The bot is built using Go and the Slack API, allowing it to interact with Slack users through custom commands.

- **cmd**: Main GO server.
- **config**: Configuration settings.
- **slack**: Slack-related integrations.

### [climatrax](./climatrax)

A climate-related project that fetches weather data via the OpenWeather API.

- **cmd**: Main GO server.
- **config**: Configuration settings.
- **handlers**: HTTP request handlers.
- **openweather**: Integration with the OpenWeather API.

### [crud](./crud)

A simple CRUD application for managing movies using Gorilla Mux (in-memory storage).

- **cmd**: Main GO server.
- **internal**:
  - **movies**: Logic for managing movie-related operations and HTTP handlers.

### [crypton](./crypton)

A blockchain implementation with transactions and proof-of-work, using BadgerDB for storage.

### [ect](./ect)

A simple email and domain checker tool (MX, SPF, DMARC records).

### [fileck](./fileck)

A simple Slack bot that uploads files to a Slack channel using the Slack API. The bot uses an OAuth token and file paths to upload files to specified channels in Slack.

- **cmd**: Main GO server.
- **config**: Configuration settings.
- **slack**: Slack-related integrations.

### [fintech](./fintech)

A financial application built with Temporal workflows, TigerBeetle, and PostgreSQL.

### [go-server](./go-server)

A Go-based HTTP server that serves static files and handles form submissions.

- **static**: Static files served by the server.

### [matrix](./matrix)

A computer systems management application, designed to simulate and manage various components of computer systems, including laptops, processors, memory, screens, keyboards, and storage devices. The system uses **gRPC** for communication and **Protocol Buffers (protobuf)** for defining data structures, enabling efficient and scalable management of hardware resources.

- **cmd**: Main GO server.
- **internal**:
  - **fns**: Functions related to various management operations, such as generating random data for system components or other utility functions.
  - **pb**: Protocol Buffers-generated files defining the data structures and services for different system components like keyboards, memory, processors, storage, and laptops.
    - **pbconnect**: Logic for establishing connections and handling communication between different protocol buffer services.
  - **serializer**: Logic for serializing and deserializing data, including support for JSON and file serialization.
  - **services**: Service implementations for handling the business logic of managing system components like laptops, processors, and storage.
  - **tmp**: Temporary or experimental code, used for testing new features or approaches.
- **proto**: Protocol Buffer definitions for services such as laptop, memory, processor, and storage, which define how data is communicated between services.

### [mchkx](./mchkx)

A simple Go program to check DNS records for a domain's mail configuration. It verifies if the domain has MX, SPF, and DMARC records and outputs the results in a CSV format.

### [pulse](./pulse)

Pulse is a real-time notification system built using Go and Kafka. This system includes a producer to send notifications and a consumer to process them.

- **cmd**:
  - **consumer**: Main GO server for consuming messages.
  - **producer**: Main GO server for producing messages.
- **internal**:
  - **errors**: Error handling logic.
  - **handlers**: HTTP request handlers.
  - **kafka**: Kafka-related integrations.
  - **middlewares**: Middlewares for various operations.
  - **models**: Database models.
  - **store**: Data storage logic.
  - **utils**: Utility functions and helpers.

### [stocks](./stocks)

A stock CRUD application with PostgreSQL database and migrations.

- **cmd**:
  - **migrate**: Database migrations.
    - **migrations**: SQL or other migration files.
- **internal**:
  - **config**: Configuration settings.
  - **handlers**: HTTP request handlers.
  - **models**: Database models.
  - **routes**: API route definitions.

### [sysdesigns](./sysdesigns)

System design implementations and study notes.

- **[shortly](./sysdesigns/shortly)**: URL shortener with MongoDB and Redis.
- **[ticketmaster](./sysdesigns/ticketmaster)**: Event ticketing platform with PostgreSQL, Redis, and sqlc.

### [ums](./ums)

A user management system with MongoDB database.

- **cmd**: Main GO server.
- **internal**:
  - **config**: Configuration settings.
  - **handlers**: HTTP request handlers.
  - **models**: Database models.
  - **routes**: API route definitions.
