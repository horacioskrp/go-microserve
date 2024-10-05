package main

import (
	"log"
	"net/http"

	"github.com/horacioskrp/common"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	httpAddr := ":8000"

	mux := http.NewServeMux()
	handler := NewHander()
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
