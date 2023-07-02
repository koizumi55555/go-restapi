package common

var (
	Scopes = []string{
		"email",
		"https://www.googleapis.com/auth/gmail.readonly",
	}
)

const (
	AuthorizationEndpoint = "https://accounts.google.com/o/oauth2/auth"
	TokenEndpoint         = "https://accounts.google.com/o/oauth2/token"
	RedirectURI           = "http://localhost:8080/auth_redirect"
)
