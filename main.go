package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"time"
)

func getMp3Files() ([]string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var mp3Files []string
	for _, file := range files {
		if !file.IsDir() && isMP3File(file.Name()) {
			mp3Files = append(mp3Files, file.Name())
		}
	}

	sort.Strings(mp3Files) // Sort the mp3 files alphabetically

	return mp3Files, nil
}

func isMP3File(filename string) bool {
	return len(filename) > 4 && filename[len(filename)-4:] == ".mp3"
}

func readPersistenceFile(filename string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func writePersistenceFile(filename string, data map[string]interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func playMP3Files() {
	mp3Files, err := getMp3Files()
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// var currentPlaying int

	// defer file.Close()

	currentPlaying := 0

	persistenceFile := "persistence.json"
	persistenceData, _ := readPersistenceFile(persistenceFile)
	if value, ok := persistenceData["currently_playing"]; ok {
		currentPlaying = value.(int)
	}

	for {

		// ensure that the currentPlaying index is within the bounds of the mp3Files
		if currentPlaying >= len(mp3Files) {
			currentPlaying = 0
		}

		fileName := mp3Files[currentPlaying]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			currentPlaying++
			continue
		}

		// Update persistence data 
		persistenceData["currently_playing"] = currentPlaying
		err2 := writePersistenceFile(persistenceFile, persistenceData)
		if err2 != nil {
			fmt.Println("Error writing persistence file:", err)
		}

		cmd := exec.Command("mpg123", file.Name())
		cmd.Run()

		currentPlaying++
	}
}

func main() {
	playMP3Files()
}
