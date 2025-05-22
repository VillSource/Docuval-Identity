package endpoints

type oidc struct {
	Auth     string
	Token    string
	UserInfo string
}

var OIDC = &oidc{
	Auth:     "/auth",
	Token:    "/token",
	UserInfo: "/userinfo",
}
