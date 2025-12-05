package transaction

type request struct {
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Reference string `json:"reference"`
}

type response struct {
	Id            string  `json:"id"`
	AccountNumber string  `json:"accountNumber"`
	UserId        string  `json:"userId"`
	Amount        string  `json:"amount"`
	Currency      string  `json:"currency"`
	Type          string  `json:"type"`
	Reference     *string `json:"reference"`
	CreatedAt     string  `json:"createdAt"`
}
