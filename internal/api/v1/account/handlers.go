package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"eaglebank/internal/api/middleware"
	"eaglebank/internal/application/account"
	"eaglebank/internal/domain/shared/helpers"

	"github.com/go-chi/chi/v5"
)

const (
	urlParamName = "accountNumber"
)

func New(userService account.Account) (*Handler, error) {
	if userService == nil {
		return nil, fmt.Errorf("user service nil when creating user handler")
	}

	return &Handler{
		accountService: userService,
	}, nil
}

type Handler struct {
	accountService account.Account
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createAccountRequest
	LoggedInUserID := r.Context().Value(middleware.UserIDKey)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if ok := h.validateCreateAccountRequest(req); !ok {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	err := h.accountService.Create(LoggedInUserID.(string), req.Name, req.AccountType)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	LoggedInUserID := r.Context().Value(middleware.UserIDKey)

	accountList, err := h.accountService.List(LoggedInUserID.(string))
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	responseJSON, err := h.outboundMappingListTranslator(accountList)
	if err != nil {
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
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	if affectedRows == 0 {
		http.Error(w, "Bank account was not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
