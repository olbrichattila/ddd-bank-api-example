package user

type createUserRequest struct {
	Name        string     `json:"name"`
	Address     addressDTO `json:"address"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
}

type createUserResponse struct {
	Id               string     `json:"id"`
	Name             string     `json:"name"`
	Address          addressDTO `json:"address"`
	PhoneNumber      string     `json:"phone_number"`
	Email            string     `json:"email"`
	CreatedTimestamp string     `json:"createdTimestamp"`
	UpdatedTimestamp string     `json:"updatedTimestamp"`
}

type addressDTO struct {
	Line1    string  `json:"line1"`
	Line2    *string `json:"line2,omitempty"`
	Line3    *string `json:"line3,omitempty"`
	Town     string  `json:"town"`
	County   string  `json:"county"`
	Postcode string  `json:"postcode"`
}
