package types

type OAuthConfig struct {
	GrantType    string // OAuth Type (default is password)
	Username     string // Username
	Password     string // Password
	ClientId     string // Client Id (default is kubesphere)
	ClientSecret string // Client Secret (default is kubesphere)
}

type OAuthToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
}
