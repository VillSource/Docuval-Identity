package models

type AuthRequestModel struct {
	GrantType       string `json:"grant_type"`
	ChallengeCode   string `json:"challenge_code"`
	AuthCode        string `json:"auth_code"`
	ChallengeMethod string `json:"challenge_method"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	Scope           string `json:"scope"`
	State           string `json:"state"`
	RedirectURI     string `json:"redirect_uri"`
	RequestTime     string `json:"request_time"`
}


