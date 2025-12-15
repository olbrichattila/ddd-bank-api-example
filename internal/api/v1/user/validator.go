package user

import "atybank/internal/shared/helpers"

func (h *Handler) validate(req createUserRequest) bool {

	if req.Name == "" || req.Email == "" {
		return false
	}

	if !helpers.IsValidEmail(req.Email) {
		h.logger.Debug("invalid email")
		return false
	}

	if !helpers.IsValidPhone(req.PhoneNumber) {
		h.logger.Debug(req.PhoneNumber)
		return false
	}

	return true
}
