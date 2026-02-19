package main

import (
	"cyber-battle/internal/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	fmt.Printf("Starting API server on :%s...\n", port)
	apiInstance := api.NewAPI()
	handler := api.NetHTTPHandler(apiInstance)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Println("API server error:", err)
	}
}
