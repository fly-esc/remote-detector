package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	_ "embed"
)

const (
	WEBSOCKET_PORT = "8080"
)

func main() {
	//log.Println("Launching server...")

	//config runtime
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)

	port := os.Getenv("PORT")
	if port == "" {
		port = WEBSOCKET_PORT
		fmt.Printf("Running with %s port\n", port)
	}

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		handleHttp(w, r)
	})
	/*http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebsocket(w, r)
	})*/

	err := http.ListenAndServe(":"+port, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server:", err)
	}
}
