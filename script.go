package main

// Libraries
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

// Get credential from json file
func getCreds(cred_file []uint8) {

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
		fmt.Print(client_id)
	}
}

// Get authorization code from the users.
func getAuthorizationCode(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Println("Go to this authurl and type in authorization code: ", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}
	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok

}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {

	// Create a token.json file to save the token.
	tokFile := "token.json"
	f, err := os.Open(tokFile)
	defer f.Close()
	if err != nil {
		panic("Unable to create a file")
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)

	// Get token by asking user for authorization code.
	if err != nil {
		tok := getAuthorizationCode(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)

}

// Check if a string exist within an array
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getField(v []*gmail.MessagePartHeader, field string) {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	fmt.Println("f", f)
}

// Main function
func main() {
	subject_keywords := []string{"Lyft", "Uber"}
	email_to_delete := []string{}
	cred_file, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	} else {
		config, err := google.ConfigFromJSON(cred_file, gmail.GmailReadonlyScope)
		if err != nil {
			panic("Error getting config")
		} else {

			// Get client and create a gmail server
			client := getClient(config)
			server, err := gmail.New(client)
			if err != nil {
				log.Fatalf("Unable to retrieve Gmail client: %v", err)
			}
			user := "me"

			// Get the headers of the message
			mes_server, err := server.Users.Messages.List(user).Do()

			for _, m := range mes_server.Messages {
				fmt.Println("m", m)
				msg, err := server.Users.Messages.Get("me", m.Id).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve message %v: %v", m.Id, err)
				}
				label_id := msg.LabelIds

				// Get only unread messages in inbox and primary category
				if stringInSlice("UNREAD", label_id) == true && stringInSlice("CATEGORY_UPDATES", label_id){
					subject := msg.Payload.Headers[19].Value
					words := strings.Split(subject, " ")
					for _, word := range words{
						if stringInSlice(word, subject_keywords) == true{
							fmt.Println("Message id", msg.Id)
							email_to_delete = append(email_to_delete, msg.Id)
						}
					}
				}
			}
			fmt.Println(email_to_delete)
		}
	}
}
