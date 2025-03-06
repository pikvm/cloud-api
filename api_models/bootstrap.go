package api_models

type BootstrapRedirect struct {
	RedirectURL string `json:"redirect_url"`
}

type BootstrapResult struct {
	AuthToken string `json:"auth_token"`
}
