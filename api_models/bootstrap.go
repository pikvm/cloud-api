package api_models

type BootstrapRedirect struct {
	BootstrapToken string `json:"bootstrap_token"`
	RedirectURL    string `json:"redirect_url"`
}

type BootstrapResult struct {
	AuthToken string `json:"auth_token"`
}
