# Gome Project

This is a simple Go project that runs an HTTP server and prints "Hello, World!" to the browser.

This application also provides a /extract-pdf endpoint that accepts PDF uploads, extracts the text, and logs it.

## How to Run

1. Open a terminal and navigate to the project directory:

    ```sh
    cd d:\git\test\gome
    ```

2. Initialize the Go module:

    ```sh
    go mod tidy
    ```

3. Run the Go program:

    ```sh
    go run main.go
    ```

4. Open a web browser and navigate to `http://localhost:8080` to see the "Hello, World!" message.

## Using the PDF Extraction Endpoint

To use the PDF extraction endpoint, send a POST request to `/extract-pdf` with the PDF file in a multipart form field named "pdf". For example:

```sh
curl -X POST -F "pdf=@path/to/your.pdf" http://localhost:8080/extract-pdf
```

This will extract the text and log it in the server console.
