package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Main function
func main() {
	cred_file, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	} else {

		// Convert credentials to json
		credentials := (string(cred_file))
		var data map[string]interface{}
		json.Unmarshal([]byte(credentials), &data)

		// Resolve nested json to get client id
		innermap, ok := data["installed"].(map[string]interface{})
		if !ok {
			panic("inner map is not a map!")
		} else {
			client_id := innermap["client_id"]
		}

	}
}
