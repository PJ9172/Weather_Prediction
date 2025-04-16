from meteostat import Point, Daily
import pandas as pd
from datetime import datetime

# 1. Define location (e.g. Pune)
location = Point(18.5204, 73.8567)  # Latitude & Longitude of Pune

# 2. Define date range
start = datetime(2022, 1, 1)
end = datetime(2023, 12, 31)

# 3. Fetch data
data = Daily(location, start, end)
data = data.fetch()

# 4. Save to CSV (optional, for future use)
data.to_csv("pune_weather.csv")

# 5. Preview data
print(data.head())
