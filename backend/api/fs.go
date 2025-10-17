package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var FILE_DIRECTORY = os.Getenv("ASSETS_DIR")

func PopulateDB() {
	fmt.Printf("Checking directory: %s\n", FILE_DIRECTORY)
	if FILE_DIRECTORY != "" {
		data, err := os.ReadFile(FILE_DIRECTORY + "/stations.json")
		if err == nil {
			json.NewDecoder(bytes.NewReader(data)).Decode(&DBStations)
		}
		data, err = os.ReadFile(FILE_DIRECTORY + "/groups.json")
		if err == nil {
			json.NewDecoder(bytes.NewReader(data)).Decode(&DBGroups)
		}
		data, err = os.ReadFile(FILE_DIRECTORY + "/questions.json")
		if err == nil {
			json.NewDecoder(bytes.NewReader(data)).Decode(&DBQuestions)
		}
		data, err = os.ReadFile(FILE_DIRECTORY + "/tokens.json")
		if err == nil {
			json.NewDecoder(bytes.NewReader(data)).Decode(&DBTokens)
			TokensInit = true
		}
		fmt.Println("DB Populated successfully")
	}
}

func StoreDB() {
	if FILE_DIRECTORY != "" {
		var sb strings.Builder
		json.NewEncoder(&sb).Encode(DBStations)
		os.WriteFile(FILE_DIRECTORY+"/stations.json", []byte(sb.String()), 0644)
		sb.Reset()
		json.NewEncoder(&sb).Encode(DBGroups)
		os.WriteFile(FILE_DIRECTORY+"/groups.json", []byte(sb.String()), 0644)
		sb.Reset()
		json.NewEncoder(&sb).Encode(DBQuestions)
		os.WriteFile(FILE_DIRECTORY+"/questions.json", []byte(sb.String()), 0644)
		sb.Reset()
		json.NewEncoder(&sb).Encode(DBTokens)
		os.WriteFile(FILE_DIRECTORY+"/tokens.json", []byte(sb.String()), 0644)
	}
}
