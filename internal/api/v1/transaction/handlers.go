package transaction

import (
	"fmt"
	"net/http"

	"eaglebank/internal/application/transaction"
)

func New(transactionService transaction.Transaction) (*Handler, error) {
	if transactionService == nil {
		return nil, fmt.Errorf("transaction handler needs transaction service to be passed, it is nil")
	}

	return &Handler{
		transactionService: transactionService,
	}, nil
}

type Handler struct {
	transactionService transaction.Transaction
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World")
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World")
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World")
}
