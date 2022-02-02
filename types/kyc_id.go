package types

type ID struct {
	FullLegalName  string `json:"full_legal_name"`
	GithubHandle   string `json:"github_handle"`
	EmailAddress   string `json:"email_address"`
	AccountAddress string `json:"account_address"`
}
