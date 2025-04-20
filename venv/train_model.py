import pandas as pd
from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split
import pickle

# Step 1: Load the weather data
df = pd.read_csv("pune_weather.csv")

# Step 2: Choose features and target
features = ['tmin', 'tmax', 'wspd', 'pres', 'prcp']
target = 'tavg'

# Step 3: Drop rows with missing values
df = df[features + [target]].dropna()

# Step 4: Split data into inputs (X) and target (y)
X = df[features]
y = df[target]

# Step 5: Split into training and test sets
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Step 6: Train the model
model = LinearRegression()
model.fit(X_train, y_train)

# Step 7: Save the trained model to file
with open("model.pkl", "wb") as f:
    pickle.dump(model, f)

print("✅ Model trained and saved as model.pkl")
