package account

type createAccountRequest struct {
	Name        string `json:"name"`
	AccountType string `json:"accountType"`
}

type accountResponse struct {
	AccountNumber    string `json:"accountNumber"`
	SortCode         string `json:"sortCode"`
	Name             string `json:"name"`
	AccountType      string `json:"accountType"`
	Balance          string `json:"balance"`
	Currency         string `json:"currency"`
	CreatedTimestamp string `json:"createdTimestamp"`
	UpdatedTimestamp string `json:"updatedTimestamp"`
}
