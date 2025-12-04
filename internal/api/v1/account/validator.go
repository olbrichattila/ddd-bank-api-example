package account

import "eaglebank/internal/domain/shared/helpers"

func (h *Handler) validateCreateAccountRequest(req createAccountRequest) bool {
	if !helpers.IsValidAccountType(req.AccountType) || req.Name == "" {
		return false
	}

	return true
}
