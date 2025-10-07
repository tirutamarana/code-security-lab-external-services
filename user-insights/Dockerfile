FROM python:3.9-slim

# Set up a working directory
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

# Expose the port the service runs on
EXPOSE 5001

# Start the application
CMD ["python", "app.py"]
