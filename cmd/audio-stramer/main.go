package main

import (
	"fmt"
	"net/http"

	"github.com/ericoliveiras/audios-streamer-go/handlers"
)

const port = 8080

func main() {
	http.HandleFunc("/musics", handlers.AudioStreamHandler)
	fmt.Printf("Server listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
