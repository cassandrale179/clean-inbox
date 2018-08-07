# clean-inbox
Delete labelled email with Gmail API. Just a practice repo for Golang. 


## Installation 
- Install [Go](https://stackoverflow.com/questions/12843063/install-go-with-brew-and-running-the-gotour)  
- Install [Dep](https://golang.github.io/dep/docs/installation.html) 
- Clone the repo under ~/Go/src/github.com/user/ 

## Useful snippets 
- Scope in Gmail API: 
```
// OAuth2 scopes used by this API.
const (
	// Read, send, delete, and manage your email
	MailGoogleComScope = "https://mail.google.com/"

	// Manage drafts and send emails
	GmailComposeScope = "https://www.googleapis.com/auth/gmail.compose"

    ... 
	// Manage your sensitive mail settings, including who can manage your mail
	GmailSettingsSharingScope = "https://www.googleapis.com/auth/gmail.settings.sharing"
)
```
- To set scope for your app (if a token is already generated with a different scope within token.json, it needs to be refreshed): 
```
config, err := google.ConfigFromJSON(credentials_file, gmail.MailGoogleComScope)
``` 

## Resources
- Gmail API Go Client: https://github.com/google/google-api-go-client/blob/master/gmail/v1/gmail-gen.go 
