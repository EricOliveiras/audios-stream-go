package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func AudioStreamHandler(w http.ResponseWriter, r *http.Request) {
	audioPath := os.Args[1]
	audioFile, err := os.Open(audioPath)
	if err != nil {
		http.Error(w, "Failed to open audio file", http.StatusInternalServerError)
		return
	}
	defer audioFile.Close()

	w.Header().Set("Content-Type", "audio/mpeg")
	_, err = io.Copy(w, audioFile)
	if err != nil {
		fmt.Println(err)
	}
}
