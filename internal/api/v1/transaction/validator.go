package transaction

import "eaglebank/internal/shared/helpers"

func (h *Handler) validateRequest(req request) bool {
	if !helpers.IsValidCurrency(req.Currency) ||
		!helpers.IsValidTransactionType(req.Type) ||
		!helpers.IsValidPaymentAmount(req.Amount) {
		return false
	}

	return true
}
