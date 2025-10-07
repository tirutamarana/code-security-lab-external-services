# 📊 User Insights API

The **User Insights API** provides analytics for **user interactions** on TechStories.  
This service tracks **post views, upvotes, and comments**, generating **engagement reports**.

---

## 📌 Features

- ✅ **Track user behavior** (views, upvotes, comments)
- 📊 **Generate analytics reports** on platform engagement
- 📡 **Fetch external user data** for enrichment
- 🛡️ **Integrate with AI-based content analysis**

---

## 🛠️ Setup & Running the Service

### **📌 Prerequisites**
- 🐍 **Python 3.9+**
- 🐳 **Docker & Docker Compose**
- 🗄️ **SQLite (or another supported database)**

### **🚀 Running with Docker**
```sh
docker build -t user-insights-api .
docker run -p 5001:5001 user-insights-api
```

### **🚀 Running Locally**
```sh
pip install -r requirements.txt
python app.py
```

### **📂 Configuration**
This service reads configuration settings from `config.yaml`. Example:
```yaml
server:
  port: 5001
  bind_address: "0.0.0.0"
```
- Modify this file to **change the database URI, API endpoints, or security settings**.

---

## 📡 API Endpoints

### **📊 Get User Activity**
**Endpoint:** `GET /get_user_activity`  
**Description:** Retrieves all interactions for a given user.

**Request Example**
```sh
curl -X GET "http://localhost:5001/get_user_activity?user_id=123"
```

**Response Example**
```json
{
  "user_id": "123",
  "activity": [
    {"post_id": 45, "interaction": "upvote"},
    {"post_id": 67, "interaction": "comment"}
  ]
}
```

---

### **📄 Generate Analytics Report**
**Endpoint:** `POST /generate_report`  
**Description:** Generates a detailed engagement report.

**Request Example**
```sh
curl -X POST "http://localhost:5001/generate_report" -d "report_type=monthly"
```

**Response Example**
```json
{"message": "Report generated successfully!"}
```

---

### **📡 External Data Enrichment**
**Endpoint:** `GET /enrich_user_data`  
**Description:** Fetches additional user details from a third-party service.

**Request Example**
```sh
curl -X GET "http://localhost:5001/enrich_user_data?user_id=123"
```

**Response Example**
```json
{
  "user_id": "123",
  "name": "John Doe",
  "engagement_score": 87
}
```

---

## 🛠️ Deployment

### **📌 Running with Docker Compose**
This service is designed to work with **TechStories** and can be deployed using `docker-compose`:
```yaml
services:
  user-insights:
    build: .
    ports:
      - "5001:5001"
    environment:
      - CONFIG_PATH=/app/config.yaml
```
Start the service:
```sh
docker-compose up -d
```

---

## 🤝 Contributing

We welcome contributions!  
- **Found a bug?** Open an issue!  
- **Want to add new analytics features?** Submit a pull request!  

---

## 📬 Contact

For any questions, reach out to the **TechStories DevOps Team**.
