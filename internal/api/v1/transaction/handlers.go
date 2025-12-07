package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"

	"eaglebank/internal/api/middleware"
	"eaglebank/internal/application/transaction"
	"eaglebank/internal/shared/helpers"
	"eaglebank/internal/shared/services"

	"github.com/go-chi/chi/v5"
	"github.com/shopspring/decimal"
)

const (
	accountNumberURLParam     = "accountNumber"
	transactionNumberURLParam = "transactionNumber"
)

type Handler struct {
	transactionService transaction.Transaction
	logger             services.Logger
}

func New(logger services.Logger, transactionService transaction.Transaction) (*Handler, error) {
	if transactionService == nil {
		return nil, fmt.Errorf("transaction handler needs transaction service to be passed, it is nil")
	}

	if logger == nil {
		return nil, fmt.Errorf("user handler requires logger, nil provided")
	}

	return &Handler{
		transactionService: transactionService,
		logger:             logger,
	}, nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	LoggedInUserId := r.Context().Value(middleware.UserIdKey)
	accountNumber := chi.URLParam(r, accountNumberURLParam)
	if !helpers.IsValidAccountNumber(accountNumber) {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if ok := h.validateRequest(req); !ok {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	decimalAmount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	err = h.transactionService.Create(decimalAmount, LoggedInUserId.(string), req.Currency, req.Type, accountNumber, &req.Reference)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	accountNumber := chi.URLParam(r, accountNumberURLParam)
	if !helpers.IsValidAccountNumber(accountNumber) {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	transactionList, err := h.transactionService.List(accountNumber)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	responseJSON, err := h.outboundMappingListTranslator(transactionList)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "User was not found", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	transactionNumber := chi.URLParam(r, transactionNumberURLParam)
	if !helpers.IsValidTransactionId(transactionNumber) {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	accountEntity, err := h.transactionService.Get(transactionNumber)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	if accountEntity == nil {
		http.Error(w, "Transaction account was not found", http.StatusNotFound)
		return
	}

	accountResponse := h.outboundMappingTranslator(accountEntity)
	accountAsJSON, err := json.Marshal(accountResponse)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	w.Write(accountAsJSON)
}
