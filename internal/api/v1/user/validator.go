package user

func (h *Handler) validate(req createUserRequest) bool {
	if req.Name == "" || req.Email == "" {
		return false
	}

	return true
}
