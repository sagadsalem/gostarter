package response

// authResponse represents an authentication response body
type AuthResponse struct {
	AccessToken string `json:"token" example:"v2.local.Gdh5kiOTyyaQ3_bNykYDeYHO21Jg2..."`
}

// newAuthResponse is a helper function to create a response body for handling authentication data
func NewAuthResponse(token string) AuthResponse {
	return AuthResponse{
		AccessToken: token,
	}
}
