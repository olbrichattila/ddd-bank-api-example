package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	userService "eaglebank/internal/application/user"

	"github.com/go-chi/chi/v5"
)

const (
	urlParamName = "userId"
)

type Handler struct {
	userService userService.User
}

func New(userService userService.User) (*Handler, error) {
	if userService == nil {
		return nil, fmt.Errorf("user handler requires user repository, nil provided")
	}

	return &Handler{
		userService: userService,
	}, nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if ok := h.validate(req); !ok {
		// I think it is 422 Unprocessable entity
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	_, err := h.userService.Create(
		req.Name,
		req.Address.Line1,
		req.Address.Line2,
		req.Address.Line3,
		req.Address.Town,
		req.Address.County,
		req.Address.Postcode,
		req.PhoneNumber,
		req.Email,
	)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, urlParamName)
	userEntity, err := h.userService.Get(userID)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	if userEntity == nil {
		http.Error(w, "User was not found", http.StatusInternalServerError)
		return
	}

	responseJSON, err := h.outboundMappingTranslator(userEntity)
	if err != nil {
		http.Error(w, "User was not found", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, urlParamName)

	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if ok := h.validate(req); !ok {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	updatedUserEntity, err := h.userService.Update(
		userID,
		req.Name,
		req.Address.Line1,
		req.Address.Line2,
		req.Address.Line3,
		req.Address.Town,
		req.Address.County,
		req.Address.Postcode,
		req.PhoneNumber,
		req.Email,
	)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	responseJSON, err := h.outboundMappingTranslator(updatedUserEntity)
	if err != nil {
		http.Error(w, "User was not found", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, urlParamName)

	affectedRows, err := h.userService.Delete(userID)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	// TODO when the delete service implements
	// checks, and errors out if user A user cannot be deleted when they are associated with a bank account
	// should return 409

	if affectedRows == 1 {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(w, "The user has been deleted")
		return
	}

	http.Error(w, "User was not found", http.StatusNotFound)
}
