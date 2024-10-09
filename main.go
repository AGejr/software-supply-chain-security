
package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Set up logging
    log.Out = os.Stdout
    log.Formatter = &logrus.JSONFormatter{}

    // Set up routing
    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Info("Received request")
        fmt.Fprintf(w, "Hello World!")
    })

    // Start server
    port := "8080"
    log.Infof("Server starting on port %s...", port)
    http.ListenAndServe(":"+port, r)
}
