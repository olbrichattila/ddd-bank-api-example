package account

import (
	"encoding/json"
	"time"

	accountDomain "eaglebank/internal/domain/account"
)

func (h *Handler) outboundMappingTranslator(accountEntity accountDomain.AccountEntity) accountResponse {
	return accountResponse{
		AccountNumber:    accountEntity.AccountNumber(),
		SortCode:         accountEntity.SortCode(),
		Name:             accountEntity.Name(),
		AccountType:      accountEntity.AccountType(),
		Balance:          accountEntity.Balance().StringFixed(2),
		Currency:         accountEntity.Currency(),
		CreatedTimestamp: accountEntity.CreatedAt().Format(time.DateTime),
		UpdatedTimestamp: accountEntity.UpdatedAt().Format(time.DateTime),
	}
}

func (h *Handler) outboundMappingListTranslator(accountEntities []accountDomain.AccountEntity) ([]byte, error) {
	responseAccounts := make([]accountResponse, len(accountEntities))

	for i, accountEntity := range accountEntities {
		responseAccounts[i] = h.outboundMappingTranslator(accountEntity)
	}

	return json.Marshal(responseAccounts)
}
