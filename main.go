package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
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

	sort.Strings(mp3Files) // Sort the mp3 files alphabetically

	return mp3Files
}

func isMP3File(filename string) bool {
	return len(filename) > 4 && filename[len(filename)-4:] == ".mp3"
}


func playMP3Files() {
	mp3Files, err := getMp3Files()
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var currentPlaying int
	for {

		// ensure that the currentPlaying index is within the bounds of the mp3Files
		if currentPlaying >= len(mp3Files) {
			currentPlaying = 0
		}

		file := mp3Files[currentPlaying]
		if file.IsDir() || !isMP3File(file.Name()) {
			currentPlaying++
			continue
		}

		// print currently playing
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