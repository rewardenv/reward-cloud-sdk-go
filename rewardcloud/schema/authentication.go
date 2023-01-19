package schema

type AuthenticationTokenRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

// AuthenticationTokenResponse defines the schema of the response
// when logging in.
type AuthenticationTokenResponse struct {
	Token string `json:"token"`
}
