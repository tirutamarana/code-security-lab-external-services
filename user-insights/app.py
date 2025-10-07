import os
import logging
import requests
import yaml
import boto3
import mysql.connector
from flask import Flask, request, jsonify

app = Flask(__name__)

s3_client = boto3.client("s3")
dynamodb_client = boto3.client("dynamodb")
sns_client = boto3.client("sns")

# Load config.yaml
with open("config.yaml", "r") as f:
    config = yaml.safe_load(f)

def get_db_connection():
    connection = mysql.connector.connect(
        host="db.techstories.com",
        user="db.example.com",
        passwd="SuperSecret123!",
        database="user_insights",
        charset='utf8mb4',
        use_pure=True,
        connection_timeout=5)
    
    return connection

@app.route('/get_user_activity', methods=['GET'])
def get_user_activity():
    user_id = request.args.get("user_id")

    conn = get_db_connection()
    cursor = conn.cursor()
    
    cursor.execute("SELECT * FROM activity_logs WHERE user_id = %s", (user_id,))
    data = cursor.fetchall()
    
    conn.close()
    return jsonify({"user_id": user_id, "activity": data})

@app.route('/analyze_user_behavior', methods=['POST'])
def analyze_user_behavior():
    user_id = request.json.get("user_id")

    response = requests.post(
        "https://analysis-service/api/analyze",
        json={"user_id": user_id},
        timeout=5
    )

    if response.status_code == 200:
        return jsonify(response.json())
    else:
        return jsonify({"error": "Failed to analyze behavior"}), 500

@app.route('/log_user_activity', methods=['POST'])
def log_user_activity():
    user_id = request.json.get("user_id")
    activity_data = request.json.get("activity_data")

    log_entry = f"User {user_id}: {activity_data}"
    
    logging.info("User activity recorded.")

    s3_client.put_object(Bucket=os.getenv("S3_BUCKET", "default-bucket"), Key=f"{user_id}.log", Body=log_entry)

    return jsonify({"message": "User activity logged"})

@app.route('/publish_user_event', methods=['POST'])
def publish_user_event():
    user_id = request.json.get("user_id")
    event_type = request.json.get("event_type")

    sns_client.publish(
        TopicArn=os.getenv("SNS_TOPIC_ARN", "arn:aws:sns:us-east-1:123456789012:UserEvents"),
        Message=f"User {user_id} performed {event_type}"
    )

    return jsonify({"message": "User event published"})

@app.route('/config', methods=['GET'])
def get_config():
    return jsonify({
        "service": "User Insights API",
        "ai_analysis_api": config["external_services"]["ai_moderation_api"]
    })

if __name__ == '__main__':
    app.run(host=config["server"]["bind_address"], port=config["server"]["port"])

