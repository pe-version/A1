# A1 - Components, APIs & Structured Services

## Service Definition

The **IoT Sensor Service** manages IoT sensor device metadata and readings within a smart home ecosystem. It provides a RESTful API to create, read, update, and delete sensor configurations including device identity, location, measurement type, and current readings. The service persists all sensor data to SQLite and enforces Bearer token authentication on all endpoints to ensure secure access to the sensor registry.

## Project Overview

**Domain:** IoT Smart Home Sensors
**Implementation Languages:** Python (FastAPI) and Go (Gin)
**Persistence:** SQLite
**Authentication:** Bearer Token
**API Specification:** [openapi.yaml](openapi.yaml)

## Stack and Tool Versions

### Python Service
| Component | Version |
|-----------|---------|
| Python | 3.11 |
| FastAPI | 0.109.0 |
| Uvicorn | 0.27.0 |
| Pydantic Settings | 2.1.0 |
| Pytest | 8.0.0 |

### Go Service
| Component | Version |
|-----------|---------|
| Go | 1.21 |
| Gin | 1.9.1 |
| go-sqlite3 | 1.14.19 |
| google/uuid | 1.5.0 |

### Infrastructure
| Tool | Version |
|------|---------|
| Docker | 20.10+ |
| Docker Compose | 3.8 (spec) |
| SQLite | 3.x |

## Project Structure

```
A1/
├── README.md                    # This file
├── openapi.yaml                 # OpenAPI 3.0 specification
├── docker-compose.yml           # Service orchestration
├── architecture.png             # Architecture diagram
├── data/
│   └── sensors.json             # Seed data for database
├── python-service/
│   ├── main.py                  # FastAPI application entry point
│   ├── config.py                # Environment configuration
│   ├── database.py              # SQLite connection and schema
│   ├── requirements.txt         # Python dependencies
│   ├── Dockerfile               # Container definition
│   ├── models/
│   │   └── sensor.py            # Pydantic data models
│   ├── repositories/
│   │   └── sensor_repository.py # Data access layer
│   ├── routers/
│   │   ├── health.py            # Health endpoint
│   │   └── sensors.py           # Sensor CRUD endpoints
│   ├── middleware/
│   │   ├── auth.py              # Bearer token validation
│   │   └── logging.py           # Correlation ID middleware
│   └── tests/
│       └── test_sensors.py      # Integration tests
└── go-service/
    ├── main.go                  # Gin application entry point
    ├── go.mod                   # Go module definition
    ├── Dockerfile               # Container definition
    ├── config/
    │   └── config.go            # Environment configuration
    ├── database/
    │   └── database.go          # SQLite connection and schema
    ├── models/
    │   └── sensor.go            # Data structs and validation
    ├── repositories/
    │   └── sensor_repository.go # Data access layer (interface)
    ├── handlers/
    │   ├── health.go            # Health endpoint handler
    │   └── sensors.go           # Sensor CRUD handlers
    ├── middleware/
    │   ├── auth.go              # Bearer token middleware
    │   └── logging.go           # Correlation ID middleware
    └── tests/
        └── sensors_test.go      # Integration tests
```

## API Endpoints

All endpoints require Bearer token authentication via the `Authorization` header.

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Service health check |
| GET | `/sensors` | List all sensors |
| GET | `/sensors/{id}` | Get sensor by ID |
| POST | `/sensors` | Create a new sensor |
| PUT | `/sensors/{id}` | Update an existing sensor |
| DELETE | `/sensors/{id}` | Delete a sensor |

### Port Mapping
- **Python Service:** http://localhost:8000
- **Go Service:** http://localhost:8080

## Running the Services

### Prerequisites
- Docker and Docker Compose installed
- Set the `API_TOKEN` environment variable (or use default)

### Quick Start
```bash
# Set your API token (optional, defaults to 'your-secret-token')
export API_TOKEN=my-secret-token

# Start both services
docker-compose up --build
```

### Start Individual Services
```bash
# Python service only
docker-compose up python-service --build

# Go service only
docker-compose up go-service --build
```

### Stop Services
```bash
docker-compose down
```

## Example Requests/Responses

### Authentication
All requests require the `Authorization: Bearer <token>` header:
```bash
curl -H "Authorization: Bearer your-secret-token" http://localhost:8000/health
```

### Health Check
```bash
curl -H "Authorization: Bearer your-secret-token" http://localhost:8000/health
```
Response:
```json
{
  "status": "ok",
  "service": "python"
}
```

### List Sensors
```bash
curl -H "Authorization: Bearer your-secret-token" http://localhost:8000/sensors
```
Response:
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
    }
  ],
  "count": 1
}
```

### Create Sensor
```bash
curl -X POST \
  -H "Authorization: Bearer your-secret-token" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Garage Door Sensor",
    "type": "contact",
    "location": "garage",
    "value": 1,
    "unit": "boolean",
    "status": "active"
  }' \
  http://localhost:8000/sensors
```
Response (201 Created):
```json
{
  "id": "sensor-007",
  "name": "Garage Door Sensor",
  "type": "contact",
  "location": "garage",
  "value": 1.0,
  "unit": "boolean",
  "status": "active",
  "last_reading": "2026-01-22T14:30:00Z"
}
```

### Update Sensor
```bash
curl -X PUT \
  -H "Authorization: Bearer your-secret-token" \
  -H "Content-Type: application/json" \
  -d '{"value": 75.0, "status": "inactive"}' \
  http://localhost:8000/sensors/sensor-001
```

### Delete Sensor
```bash
curl -X DELETE \
  -H "Authorization: Bearer your-secret-token" \
  http://localhost:8000/sensors/sensor-001
```
Response: 204 No Content

### Unauthorized Request (No Token)
```bash
curl http://localhost:8000/sensors
```
Response (401):
```json
{
  "detail": "Not authenticated"
}
```

## Running Tests

### Python Tests
```bash
cd python-service
pip install -r requirements.txt
pytest tests/ -v
```

### Go Tests
```bash
cd go-service
go test ./tests/ -v
```

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────────────┐
│                           EXTERNAL CLIENTS                               │
│                     (curl, Postman, Web Apps, etc.)                     │
└────────────────────────────────┬────────────────────────────────────────┘
                                 │
                                 │ HTTPS/HTTP
                                 │ Authorization: Bearer <token>
                                 ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                        SECURITY BOUNDARY                                 │
│  ┌───────────────────────────────────────────────────────────────────┐  │
│  │                    Bearer Token Validation                         │  │
│  │                  (Middleware - All Routes)                         │  │
│  └───────────────────────────────────────────────────────────────────┘  │
└────────────────────────────────┬────────────────────────────────────────┘
                                 │
                                 │ Validated Requests
                                 ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                         IoT SENSOR SERVICE                               │
│                                                                          │
│  ┌─────────────────────────────────────────────────────────────────┐    │
│  │                    LOGGING MIDDLEWARE                            │    │
│  │              (Correlation ID, Request/Response Logging)          │    │
│  └─────────────────────────────────────────────────────────────────┘    │
│                                 │                                        │
│  ┌──────────────────────────────┼──────────────────────────────────┐    │
│  │                         API LAYER                                │    │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐   │    │
│  │  │ GET /health  │  │ GET /sensors │  │ POST/PUT/DELETE      │   │    │
│  │  │              │  │ GET /:id     │  │ /sensors/:id         │   │    │
│  │  └──────────────┘  └──────────────┘  └──────────────────────┘   │    │
│  └──────────────────────────────┼──────────────────────────────────┘    │
│                                 │                                        │
│  ┌──────────────────────────────┼──────────────────────────────────┐    │
│  │                      DEPENDENCY INJECTION                        │    │
│  │           (Python: Depends() / Go: Interface-based)              │    │
│  └──────────────────────────────┼──────────────────────────────────┘    │
│                                 │                                        │
│  ┌──────────────────────────────┼──────────────────────────────────┐    │
│  │                      REPOSITORY LAYER                            │    │
│  │              (SensorRepository - Data Access)                    │    │
│  └──────────────────────────────┼──────────────────────────────────┘    │
│                                 │                                        │
└─────────────────────────────────┼────────────────────────────────────────┘
                                  │
                                  │ SQL Queries
                                  ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                          DATA STORAGE                                    │
│                                                                          │
│  ┌───────────────────────────────────────────────────────────────────┐  │
│  │                        SQLite Database                             │  │
│  │                      /app/data/sensors.db                          │  │
│  │                                                                    │  │
│  │  ┌─────────────────────────────────────────────────────────────┐  │  │
│  │  │  sensors TABLE                                               │  │  │
│  │  │  - id, name, type, location, value, unit, status,            │  │  │
│  │  │    last_reading, created_at, updated_at                      │  │  │
│  │  └─────────────────────────────────────────────────────────────┘  │  │
│  └───────────────────────────────────────────────────────────────────┘  │
│                                                                          │
│  ┌───────────────────────────────────────────────────────────────────┐  │
│  │  Docker Volume Mount: ./data:/app/data (persistent storage)       │  │
│  └───────────────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## How Each Framework Handles DI and Configuration

### Python (FastAPI)

**Dependency Injection:**
FastAPI uses a declarative dependency injection system via the `Depends()` function. Dependencies are declared as function parameters and resolved automatically at request time.

```python
# Example: Repository dependency injection
@router.get("/sensors")
def list_sensors(
    repo: SensorRepository = Depends(get_sensor_repository),
    _: str = Depends(verify_token),  # Auth check
):
    return repo.get_all()
```

Key characteristics:
- Dependencies are functions that yield or return values
- Can be nested (dependencies can have their own dependencies)
- Scoped per-request by default
- Supports async dependencies
- Built-in support for security schemes (HTTPBearer, OAuth2)

**Configuration:**
Uses `pydantic-settings` for type-safe configuration with automatic environment variable loading:

```python
class Settings(BaseSettings):
    api_token: str  # Required
    database_path: str = "/app/data/sensors.db"  # Default

    class Config:
        env_file = ".env"
```

### Go (Gin)

**Dependency Injection:**
Go uses interface-based dependency injection through constructor functions. This is a manual but explicit pattern.

```go
// Interface defines the contract
type SensorRepository interface {
    GetAll() ([]Sensor, error)
    Create(sensor *SensorCreate) (*Sensor, error)
}

// Concrete implementation
type SQLiteSensorRepository struct {
    db *sql.DB
}

// Handler receives repository via constructor
func NewSensorHandler(repo SensorRepository) *SensorHandler {
    return &SensorHandler{repo: repo}
}
```

Key characteristics:
- Interfaces are implicit (no `implements` keyword)
- Dependencies wired explicitly in `main()`
- Easy to mock for testing
- Compile-time type safety
- No runtime reflection overhead

**Configuration:**
Uses environment variables with explicit loading:

```go
type Config struct {
    Port         int
    DatabasePath string
    APIToken     string
}

func Load() (*Config, error) {
    apiToken := os.Getenv("API_TOKEN")
    if apiToken == "" {
        return nil, fmt.Errorf("API_TOKEN required")
    }
    return &Config{APIToken: apiToken}, nil
}
```

---

## Trade-offs Observed

| Aspect | Python (FastAPI) | Go (Gin) |
|--------|------------------|----------|
| **DI Simplicity** | Built-in, declarative | Manual, explicit |
| **Type Safety** | Runtime validation via Pydantic | Compile-time checking |
| **Configuration** | Pydantic auto-loads .env files | Manual env parsing |
| **Error Handling** | Exceptions, auto-converted to HTTP | Explicit error returns |
| **Testing** | TestClient with fixtures | httptest package |
| **Boilerplate** | Less code for same functionality | More verbose but explicit |
| **Performance** | Fast (async-capable) | Very fast (compiled binary) |
| **Binary Size** | ~150MB Docker image | ~15MB Docker image |
| **Learning Curve** | Gentle, Pythonic | Steeper, different paradigms |

### Key Observations

1. **DI Trade-off:** FastAPI's `Depends()` is elegant but "magical" - Go's explicit wiring is more verbose but clearer about what's happening.

2. **Configuration Trade-off:** Pydantic settings validate at startup and fail fast. Go's manual approach requires explicit validation but has no hidden behavior.

3. **Error Handling Trade-off:** FastAPI exceptions auto-convert to HTTP responses (less code), but Go's explicit error returns make error paths obvious in the code.

4. **Validation Trade-off:** Pydantic models validate automatically. Go requires manual validation but gives full control over error messages.

5. **Middleware Trade-off:** Both frameworks handle middleware similarly, but Gin requires explicit `c.Abort()` to stop the chain.

---

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

---

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `API_TOKEN` | Yes | `your-secret-token` | Bearer token for authentication |
| `DATABASE_PATH` | No | `sensors-python.db` / `sensors-go.db` | SQLite database file path (separate per service) |
| `SEED_DATA_PATH` | No | `/app/data/sensors.json` | JSON file to seed database |
| `LOG_LEVEL` | No | `INFO` | Logging level |
| `LOG_FORMAT` | No | `json` | Log format (json or text) |
| `PORT` | No | 8000/8080 | Service port |

**Note:** Each service maintains its own separate SQLite database (`sensors-python.db` for Python, `sensors-go.db` for Go) to allow both services to run simultaneously without database conflicts.

---

## Validation Rules

| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| name | string | Yes | 1-100 characters |
| type | string | Yes | One of: temperature, motion, humidity, light, air_quality, co2, contact, pressure |
| location | string | Yes | 1-100 characters |
| value | number | Yes | Any numeric value |
| unit | string | Yes | 1-50 characters |
| status | string | Yes | One of: active, inactive, error |

---

## Appendix: Reflection (from A0)

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
