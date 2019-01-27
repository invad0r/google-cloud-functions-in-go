package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/invad0r/google-cloud-functions-in-go/image_resizer"
)

// server for local testing
func main() {
	port := flag.Int("p", 8080, "server port")
	mux := http.NewServeMux()
	mux.HandleFunc("/ResizeImage", image_resizer.ResizeImage)
	fmt.Printf("Starting local server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
