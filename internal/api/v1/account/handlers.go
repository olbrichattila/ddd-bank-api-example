package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"eaglebank/internal/api/middleware"
	"eaglebank/internal/application/account"
	"eaglebank/internal/shared/helpers"
	"eaglebank/internal/shared/services"

	"github.com/go-chi/chi/v5"
)

const (
	urlParamName = "accountNumber"
)

type Handler struct {
	accountService account.Account
	logger         services.Logger
}

func New(logger services.Logger, userService account.Account) (*Handler, error) {
	if userService == nil {
		return nil, fmt.Errorf("user service nil when creating user handler")
	}

	if logger == nil {
		return nil, fmt.Errorf("user handler requires logger, nil provided")
	}

	return &Handler{
		accountService: userService,
		logger:         logger,
	}, nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createAccountRequest
	LoggedInUserId := r.Context().Value(middleware.UserIdKey)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if ok := h.validateCreateAccountRequest(req); !ok {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	err := h.accountService.Create(LoggedInUserId.(string), req.Name, req.AccountType)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	LoggedInUserId := r.Context().Value(middleware.UserIdKey)

	accountList, err := h.accountService.List(LoggedInUserId.(string))
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	responseJSON, err := h.outboundMappingListTranslator(accountList)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "User was not found", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	accountNumber := chi.URLParam(r, urlParamName)
	if !helpers.IsValidAccountNumber(accountNumber) {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	accountEntity, err := h.accountService.Get(accountNumber)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	if accountEntity == nil {
		http.Error(w, "Bank account was not found", http.StatusNotFound)
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

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	accountNumber := chi.URLParam(r, urlParamName)
	if !helpers.IsValidAccountNumber(accountNumber) {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	var req createAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if ok := h.validateCreateAccountRequest(req); !ok {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	modifiedAccountEntity, err := h.accountService.Update(accountNumber, req.Name, req.AccountType)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	if modifiedAccountEntity == nil {
		http.Error(w, "Bank account was not found", http.StatusNotFound)
		return
	}

	accountResponse := h.outboundMappingTranslator(modifiedAccountEntity)
	accountAsJSON, err := json.Marshal(accountResponse)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	w.Write(accountAsJSON)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	accountNumber := chi.URLParam(r, urlParamName)
	if !helpers.IsValidAccountNumber(accountNumber) {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	affectedRows, err := h.accountService.Delete(accountNumber)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	if affectedRows == 0 {
		http.Error(w, "Bank account was not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
