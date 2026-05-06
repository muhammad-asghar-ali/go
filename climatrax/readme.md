# Climatrax: Weather Tracker API

## Overview

A weather data fetching application that integrates with the OpenWeather API to provide current weather information.

## Prerequisites

Before running the Weather Tracker API, ensure you have the following installed on your system:

### 1. Go

- **Version**: Go 1.23 or later
- **Installation**: [Download Go](https://golang.org/dl/)

### 2. OpenWeather API Key

- **Description**: Obtain an API key from [OpenWeather](https://openweathermap.org/api).
- **Usage**: Add the key to a `.env` file in the project root:
  ```env
  API_KEY=your_api_key_here
  ```

## Project Structure

```
climatrax/
├── cmd/          # Main server entry point
├── config/       # Configuration settings
├── handlers/     # HTTP request handlers
│   ├── handler.go
│   └── utils.go
└── openweather/  # OpenWeather API integration
```

## Running the Application

1. Set up your `.env` file with the API key
2. Run the server:
   ```bash
   go run cmd/main.go
   ```
