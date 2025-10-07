# 🛠️ External Services for TechStories

Welcome to the **External Services** repository! This repo contains two standalone microservices that integrate with the **TechStories** platform:  
- 📊 `python-user-insights` – A **user analytics API** that tracks activity on the TechStories platform.  
- 🖼️ `go-image-processing` – A **content moderation & image transformation service** for user-uploaded images.  

Each service is **containerized** and can be run independently.

---

## 📂 Repository Structure

```
external-services/
│── python-user-insights/         # Python-based analytics API
│   ├── app.py                    # Main Flask application
│   ├── requirements.txt           # Python dependencies
│   ├── Dockerfile                 # Container setup
│   ├── config.yaml                # Example config file
│   ├── README.md                  # Explanation of the Python service
│   └── tests/                      # Simple unit tests
│       ├── test_api.py             # Tests for API endpoints
│       └── test_security.py        # Security-related tests
│
│── go-image-processing/           # Go-based image processing API
│   ├── main.go                     # Main Go application
│   ├── Dockerfile                  # Container setup
│   ├── config.yaml                 # Example config file
│   ├── README.md                   # Explanation of the Go service
│   └── tests/                       # Simple unit tests
│       ├── image_test.go           # Tests for image processing functions
│       ├── security_test.go        # Security-related tests
│
│── .github/workflows/              # CI/CD pipeline for security & builds
│   ├── docker.yml                    # Builds & pushes Docker images
│
│── docker-compose.yml              # Runs both services in containers
│── README.md                       # Overview of the repo
│── .gitignore                       # Ignore build files & sensitive data
```

---

## 🛠️ Setup & Running the Services

### **📌 Prerequisites**
- 🐍 **Python 3.9+**  
- 🦫 **Go 1.18+**  
- 🐳 **Docker & Docker Compose**  

### **🚀 Running Both Services with Docker**
```sh
docker-compose up --build
```
This command will:

-Build & start the Python user-insights service on port 5001.

-Build & start the Go image-processing service on port 8080.

## 📊 Python Service: `user-insights-api`
**Description:**  
The `user-insights-api` service tracks **user activity** (comments, upvotes, post views) and generates analytics reports.

## 🖼️ Go Service: `image-processing-api`
**Description:**  
The `image-processing-api` allows users to **upload images** and **send them for AI-based moderation**.



