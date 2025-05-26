# Weather Tracker

A simple Go application that fetches real-time weather data using OpenWeatherMap API.

## Quick Start

1. Create `.apiConfig` file:
```json
{
    "openWeatherMapAPIKey": "your-api-key-here"
}
```

2. Run the server:
```bash
go run main.go
```

3. Access weather data:
- `http://localhost:8080/weather/{cityname}`
- Example: `http://localhost:8080/weather/london`

## API Endpoints

- `GET /weather/{cityname}` - Get weather data
- `GET /hello` - Test endpoint

## Response Example
```json
{
    "name": "London",
    "main": {
        "temp": 293.15
    }
}
```

## Requirements
- Go 1.x
- OpenWeatherMap API key 