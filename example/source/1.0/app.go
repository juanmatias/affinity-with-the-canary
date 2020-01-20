package main

import (
	"fmt"
	"net/http"
)

const version string = "1.0"

func getFrontpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Congratulations! Version %s of your application is running on Kubernetes.", version)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getHeaders(w http.ResponseWriter, r *http.Request) {
        // Loop over header names
        for name, values := range r.Header {
            // Loop over all values for the name.
            for _, value := range values {
                fmt.Fprintf(w, name, value)
            }
        }
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", version)
}

func main() {
	http.HandleFunc("/", getFrontpage)
	http.HandleFunc("/health", health)
	http.HandleFunc("/headers", getHeaders)
	http.HandleFunc("/version", getVersion)
	http.ListenAndServe(":8080", nil)
}
