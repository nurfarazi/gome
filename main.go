package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    
    "github.com/pdfcpu/pdfcpu/pkg/api"
    httpSwagger "github.com/swaggo/http-swagger"
)

// @title PDF Text Extraction API
// @version 1.0
// @description API for extracting text from PDF files
// @host localhost:8080
// @BasePath /

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

// @Summary Extract text from PDF
// @Description Uploads a PDF file and extracts its text content
// @Accept multipart/form-data
// @Produce json
// @Param pdf formData file true "PDF file to process"
// @Success 200 {string} string "PDF text extracted and logged successfully"
// @Router /extract-pdf [post]
func extractPDFHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse the multipart form
    err := r.ParseMultipartForm(10 << 20) // 10 MB max
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }

    file, _, err := r.FormFile("pdf")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Create a temporary file
    tempFile, err := os.CreateTemp("", "upload-*.pdf")
    if err != nil {
        http.Error(w, "Error creating temp file", http.StatusInternalServerError)
        return
    }
    defer os.Remove(tempFile.Name())
    defer tempFile.Close()

    // Copy uploaded file to temp file
    _, err = io.Copy(tempFile, file)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }

    // Extract text from PDF
    text, err := api.ExtractText(tempFile.Name(), nil)
    if err != nil {
        http.Error(w, "Error extracting text from PDF", http.StatusInternalServerError)
        return
    }

    // Log the extracted text
    log.Printf("Extracted text from PDF: %s", text)

    fmt.Fprintf(w, "PDF text extracted and logged successfully")
}

func main() {
    // Swagger documentation endpoint
    http.HandleFunc("/swagger/*", httpSwagger.Handler(
        httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
    ))

    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/extract-pdf", extractPDFHandler)
    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)
}
