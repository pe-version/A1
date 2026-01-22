from fastapi import FastAPI
import json
from pathlib import Path
from pydantic import BaseModel
from typing import Union

app = FastAPI(title="IoT Sensor Service - Python")

DATA_FILE = Path("/app/data/sensors.json")


class Sensor(BaseModel):
    id: str
    name: str
    type: str
    location: str
    value: Union[float, bool, int]
    unit: str
    status: str
    last_reading: str


def load_sensors():
    if DATA_FILE.exists():
        with open(DATA_FILE) as f:
            return json.load(f)
    return []


@app.get("/health")
def health():
    return {"status": "ok", "service": "python"}


@app.get("/items")
def get_items():
    sensors = load_sensors()
    return {"sensors": sensors, "count": len(sensors)}
