package api_models

type WhoamiResult struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	User        WhoamiUserInfo `json:"user"`
	DefaultFqdn string         `json:"default_fqdn"`
}

type WhoamiUserInfo struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
