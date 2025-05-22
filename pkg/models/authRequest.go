package models

type AuthRequestModel struct {
	ResponseType    string `json:"response_type"`
	ChallengeCode   string `json:"code_challenge"`
	AuthCode        string `json:"code"`
	ChallengeMethod string `json:"code_challenge_method"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	Scope           string `json:"scope"`
	State           string `json:"state"`
	RedirectURI     string `json:"redirect_uri"`
	RequestTime     int64    `json:"request_time"`
}
