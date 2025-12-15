package transaction

import (
	"encoding/json"
	"time"

	transactionDomain "atybank/internal/domain/transaction"
)

func (h *Handler) outboundMappingTranslator(entity transactionDomain.TransactionEntity) response {
	return response{
		Id:            entity.Id(),
		AccountNumber: entity.AccountNumber(),
		UserId:        entity.UserId(),
		Amount:        entity.Amount().StringFixed(2),
		Currency:      entity.Currency(),
		Type:          entity.Type(),
		Reference:     entity.Reference(),
		CreatedAt:     entity.CreatedAt().Format(time.DateTime),
	}
}

func (h *Handler) outboundMappingListTranslator(entities []transactionDomain.TransactionEntity) ([]byte, error) {
	response := make([]response, len(entities))

	for i, entity := range entities {
		response[i] = h.outboundMappingTranslator(entity)
	}

	return json.Marshal(response)
}
