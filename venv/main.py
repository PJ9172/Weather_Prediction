from fastapi import FastAPI
from pydantic import BaseModel
import pickle
import numpy as np

# Load the trained model
with open("model.pkl", "rb") as f:
    model = pickle.load(f)

# Define FastAPI app
app = FastAPI()

# Define input structure
class WeatherFeatures(BaseModel):
    tmin: float
    tmax: float
    wspd: float
    pres: float
    prcp: float

# Prediction endpoint
@app.post("/predict")
def predict_temperature(data: WeatherFeatures):
    features = np.array([[data.tmin, data.tmax, data.wspd, data.pres, data.prcp]])
    prediction = model.predict(features)
    return {
        "predicted_tavg": round(prediction[0], 2)
    }