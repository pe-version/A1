# A0 - Environment & Baseline Service (Polyglot Setup)

## Project Overview

**Domain:** IoT Smart Home Sensors
**Implementation Languages:** Python (FastAPI) and Go (Gin)
**HTTP Client:** curl

## Stack and Tool Versions

### Python Service
| Component | Version |
|-----------|---------|
| Python    | 3.11    |
| FastAPI   | 0.109.0 |
| Uvicorn   | 0.27.0  |

### Go Service
| Component | Version |
|-----------|---------|
| Go        | 1.21    |
| Gin       | 1.9.1   |

### Infrastructure
| Tool           | Version    |
|----------------|------------|
| Docker         | 20.10+     |
| Docker Compose | 3.8 (spec) |

## Project Structure

```
microservices-jan-2026/
└── A0/
    ├── README.md              # This file
    ├── docker-compose.yml     # Service orchestration
    ├── data/
    │   └── sensors.json       # Sample IoT sensor data
    ├── python-service/
    │   ├── main.py            # FastAPI application
    │   ├── requirements.txt   # Python dependencies
    │   └── Dockerfile         # Container definition
    └── go-service/
        ├── main.go            # Gin application
        ├── go.mod             # Go module definition
        └── Dockerfile         # Container definition
```

## Running the Services

### Quick Start (Pre-built Images from Docker Hub)
```bash
docker-compose up
```
This pulls pre-built images from Docker Hub:
- `hiphophippo/iot-python-service:a0`
- `hiphophippo/iot-go-service:a0`

> **Note:** Images are tagged with `:a0` to ensure graders see exactly the version submitted for this assignment. The `:a0` tag is immutable, whereas `:latest` may change as work continues on future assignments.

### Build from Source
```bash
docker-compose up --build
```

### Start Individual Services
```bash
# Python service only
docker-compose up python-service

# Go service only
docker-compose up go-service
```

### Stop Services
```bash
docker-compose down
```

## API Endpoints

Both services expose identical endpoints:

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/health` | GET | Health check returning service status |
| `/items` | GET | Returns all IoT sensor data |

### Port Mapping
- **Python Service:** http://localhost:8000
- **Go Service:** http://localhost:8080

## Sample API Responses

### Health Endpoint

**Python Service** (`curl http://localhost:8000/health`):
```json
{
  "status": "ok",
  "service": "python"
}
```

**Go Service** (`curl http://localhost:8080/health`):
```json
{
  "status": "ok",
  "service": "go"
}
```

### Items Endpoint

**Python Service** (`curl http://localhost:8000/items`):
```json
{
  "sensors": [
    {
      "id": "sensor-001",
      "name": "Living Room Thermostat",
      "type": "temperature",
      "location": "living_room",
      "value": 72.5,
      "unit": "fahrenheit",
      "status": "active",
      "last_reading": "2026-01-15T10:30:00Z"
    },
    ...
  ],
  "count": 6
}
```

**Go Service** (`curl http://localhost:8080/items`):
```json
{
  "sensors": [
    {
      "id": "sensor-001",
      "name": "Living Room Thermostat",
      "type": "temperature",
      "location": "living_room",
      "value": 72.5,
      "unit": "fahrenheit",
      "status": "active",
      "last_reading": "2026-01-15T10:30:00Z"
    },
    ...
  ],
  "count": 6
}
```

## Sample Dataset

The `data/sensors.json` file contains 6 IoT sensors representing a smart home setup:

| Sensor ID | Name | Type | Location | Value |
|-----------|------|------|----------|-------|
| sensor-001 | Living Room Thermostat | temperature | living_room | 72.5°F |
| sensor-002 | Front Door Motion | motion | front_entrance | false |
| sensor-003 | Kitchen Humidity | humidity | kitchen | 45.2% |
| sensor-004 | Bedroom Light | light | master_bedroom | 0% |
| sensor-005 | Backyard Air Quality | air_quality | backyard | AQI 42 |
| sensor-006 | Office CO2 Monitor | co2 | home_office | 620 ppm |

## Smoke Test Commands

```bash
# Test Python service health
curl -s http://localhost:8000/health | jq .

# Test Python service items
curl -s http://localhost:8000/items | jq .

# Test Go service health
curl -s http://localhost:8080/health | jq .

# Test Go service items
curl -s http://localhost:8080/items | jq .
```

## Screenshots

![Both services running in Docker](go-and-python-microservices-running.png)

---

## Reflection

### Why I Chose IoT Smart Home Sensors

In short, I chose IoT Smart Home Sensors because they sounded fun and potentially useful. Last term, one of my classmates had a sensor-based project and I thought that it might be interesting to try when it was offered in this course. Moreover, I might actually try to implement some version of this in my home with actual sensors and integrate it into other personal projects. My undergraduate major was electrical engineering, and while I haven't worked in that industry since graduating, I think it's time to dip my toes back in.

### Why I Chose Python and Go

I relied on Claude's recommendations in the process of selecting my two languages to compare, which ended up being **Python** and **Go**. I was inclined toward **Python** going into this process anyway as I was familiar with it.

I chose **Python (FastAPI)** because it's so well-documented, easy-to-use, and widespread. I would say Python is my go-to language for many use cases anyway, so it seemed natural. Also, since I've been "vibe coding," I'm pretty confident reviewing Python.

I chose **Go with Gin** at Claude's recommendation for the contrast with Python, the industry applicability, and at least a little bit my own desire to do more with it. I have only used Go very slightly in a previous job, and it seems like a language I should be a bit more familiar with.

### What Worked vs. What Confused Me

**What Worked:**
Most of this worked pretty well. The initial files were pretty boilerplate (I reviewed them, asked for a few changes that matched my home's setup, etc.), and I learned some things about Go syntax (for example, := as variable declaration AND assignment as opposed to = for just assignment). I also set up Postman to have available but decided not to use it for this small use case, as it seemed unnecessary.

**What Confused Me:**

Some of the Go syntax was new to me: I needed to learn how it worked and what it meant.
I also decided to add a Sensor class to the Python after reflection for greater robustness even though this light application doesn't really necessitate it.
Another thing I learned is that when adding tags to the Docker containers, "latest" is unstable and bad practice for any sort of submission (and I would imagine especially production), as I will continue to modify things for future assignments (and, if in a commercial application, future versions). I suppose it didn't occur to me that "latest" is what the AI would initially recommend. I continue to learn to question more and more things, even as I thought I questioned some things well enough. There are many little lessons buried in these tasks.