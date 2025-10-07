# 🖼️ Image Processing API

The **Image Processing API** is a microservice that provides **image upload, transformation, and AI-based moderation** capabilities.  
This service is designed to handle **user profile pictures, post images, and other media assets** for TechStories.

---

## 📌 Features

- ✅ **Upload images** and store them for processing.
- 🔄 **Apply transformations** (grayscale, resize, rotate).
- 🛡️ **AI-based content moderation** to ensure compliance with platform rules.
- 📡 **Integration with external AI analysis services** for additional safety checks.

---

## 🛠️ Setup & Running the Service

### **📌 Prerequisites**
- 🦫 **Go 1.18+**
- 🐳 **Docker & Docker Compose**
- 🌍 **An internet connection for AI moderation API calls**

### **🚀 Running with Docker**
```sh
docker build -t image-processing-api .
docker run -p 8080:8080 image-processing-api
```

### **🚀 Running Locally**
```sh
go run main.go
```

### **📂 Configuration**
The service is configured using `config.yaml`. Example:
```yaml
server:
  port: 8080
  bind_address: "0.0.0.0"

storage:
  upload_dir: "/tmp"
```
- Modify this file to **change the upload directory, API endpoints, or security settings**.

---

## 📡 API Endpoints

### **📤 Upload an Image**
**Endpoint:** `POST /upload`  
**Description:** Accepts image files and applies transformations.

**Request Example (File Upload via cURL)**
```sh
curl -X POST -F "image=@path/to/image.jpg" "http://localhost:8080/upload?transform=grayscale"
```

**Response Example**
```json
{"message": "Image processed successfully"}
```

---

### **🔍 Request AI Content Moderation**
**Endpoint:** `GET /analyze`  
**Description:** Sends an image URL to the AI moderation service.

**Request Example**
```sh
curl -X GET "http://localhost:8080/analyze?image_url=https://example.com/image.jpg"
```

**Response Example**
```json
{"message": "AI content moderation completed"}
```

---

### **📡 External AI Analysis via gRPC**
**Endpoint:** `GET /grpc-analyze`  
**Description:** Calls an AI moderation service via gRPC.

**Request Example**
```sh
curl -X GET "http://localhost:8080/grpc-analyze"
```

**Response Example**
```json
{"message": "AI Analysis request sent!"}
```

---

## 🛠️ Deployment

### **📌 Running with Docker Compose**
This service is designed to work with **TechStories** and can be deployed using `docker-compose`:
```yaml
services:
  image-processing:
    build: .
    ports:
      - "8080:8080"
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
- **Want to add new image transformations?** Submit a pull request!

---

## 📬 Contact

For any questions, reach out to the **TechStories DevOps Team**.
