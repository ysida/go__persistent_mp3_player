package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func getMp3Files() []string {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	var mp3Files []string
	for _, file := range files {
		if !file.IsDir() && isMP3File(file.Name()) {
			mp3Files = append(mp3Files, file.Name())
		}
	}
	return mp3Files
}

func isMP3File(filename string) bool {
	return len(filename) > 4 && filename[len(filename)-4:] == ".mp3"
}


func playMP3Files() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var currentPlaying int
	for {
		if currentPlaying >= len(files) {
			currentPlaying = 0
		}

		file := files[currentPlaying]
		if file.IsDir() || !isMP3File(file.Name()) {
			currentPlaying++
			continue
		}

		fmt.Println("Playing:", file.Name())
		cmd := exec.Command("mpg123", file.Name())
		cmd.Run()

		// Persist current playing location every 5 seconds
		time.Sleep(5 * time.Second)
		fmt.Println("Persisting current playing location...")

		currentPlaying++
	}
}


func main() {
	playMP3Files()
}