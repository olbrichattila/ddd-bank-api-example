package account

import "eaglebank/internal/shared/helpers"

func (h *Handler) validateCreateAccountRequest(req createAccountRequest) bool {
	if !helpers.IsValidAccountType(req.AccountType) || req.Name == "" {
		return false
	}

	return true
}
