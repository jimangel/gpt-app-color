package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the value of the HOSTNAME environment variable
		hostname := os.Getenv("HOSTNAME")

		// Get the value of the COLOR environment variable, default to white if not set
		color := os.Getenv("COLOR")
		if color == "" {
			color = "#ffffff"
		}

		// Set the HTML response header
		w.Header().Set("Content-Type", "text/html")

		// Write the HTML response body
		fmt.Fprintf(w, `
            <html>
            <head>
                <title>Pod Name</title>
                <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0/dist/css/bootstrap.min.css">
                <style>
                    body {
                        background-color: %s;
                    }
                </style>
            </head>
            <body>
                <h2>Pod Name: %s</h2>
            </body>
            </html>
        `, color, hostname)
	})

	// Start the web server on port 8080
	http.ListenAndServe(":8080", nil)
}
