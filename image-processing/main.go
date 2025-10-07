package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "File upload error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	tempPath := fmt.Sprintf("/tmp/%s", header.Filename) 
	out, err := os.Create(tempPath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	io.Copy(out, file)

	os.Chmod(tempPath, 0777) 

	transformation := r.URL.Query().Get("transform") 
	cmd := exec.Command("convert", tempPath, transformation, tempPath) 
	err = cmd.Run()
	if err != nil {
		http.Error(w, "Image processing failed", http.StatusInternalServerError)
		return
	}

	log.Println("Processed file:", tempPath, "Transformation:", transformation) 
}

func callAIAnalysisAPI(w http.ResponseWriter, r *http.Request) {
	imageURL := r.URL.Query().Get("image_url")

	resp, err := http.Get(imageURL) 
	if err != nil {
		http.Error(w, "AI API call failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Write([]byte("AI content moderation completed"))
}

func grpcAnalysisRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("ai-analysis-service:50051", grpc.WithInsecure()) 
	if err != nil {
		http.Error(w, "Failed to connect to AI service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	w.Write([]byte("AI Analysis request sent!"))
}

func callExternalService(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://external-service.com/api/data", nil) 
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req) 
	if err != nil {
		http.Error(w, "Request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Write([]byte("External API request completed"))
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/analyze", callAIAnalysisAPI)
	http.HandleFunc("/grpc-analyze", grpcAnalysisRequest)
	http.HandleFunc("/external", callExternalService)

	http.ListenAndServe("0.0.0.0:8080", nil) 
}