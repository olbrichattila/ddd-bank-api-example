package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	userService "eaglebank/internal/application/user"
	configRepository "eaglebank/internal/infrastructure/config"
	"eaglebank/internal/shared/helpers"
	"eaglebank/internal/shared/services"

	"github.com/go-chi/chi/v5"
)

const (
	urlParamName = "userId"
)

type Handler struct {
	userService userService.User
	logger      services.Logger
	cfg         configRepository.Config
}

func New(cfg configRepository.Config, logger services.Logger, userService userService.User) (*Handler, error) {
	if cfg == nil {
		return nil, fmt.Errorf("user handler requires config repository, nil provided")
	}

	if userService == nil {
		return nil, fmt.Errorf("user handler requires user repository, nil provided")
	}

	if logger == nil {
		return nil, fmt.Errorf("user handler requires logger, nil provided")
	}

	return &Handler{
		userService: userService,
		cfg:         cfg,
		logger:      logger,
	}, nil
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO need hashed password, it is only for testing, trying this test API
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if !helpers.IsValidEmail(req.Email) {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	userEntity, err := h.userService.GetByEmail(req.Email)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	if userEntity == nil {
		http.Error(w, "unathorized", http.StatusUnauthorized)
		return
	}

	token, err := helpers.CreateToken(userEntity.Id().AsString(), []byte(h.cfg.GetJWTSecret()))
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, token)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(err.Error())
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
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, urlParamName)
	userEntity, err := h.userService.Get(userId)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	if userEntity == nil {
		http.Error(w, "User was not found", http.StatusInternalServerError)
		return
	}

	responseJSON, err := h.outboundMappingTranslator(userEntity)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "User was not found", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, urlParamName)

	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if ok := h.validate(req); !ok {
		http.Error(w, "Invalid details supplied", http.StatusBadRequest)
		return
	}

	updatedUserEntity, err := h.userService.Update(
		userId,
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
		h.logger.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	responseJSON, err := h.outboundMappingTranslator(updatedUserEntity)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "User was not found", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, urlParamName)

	affectedRows, err := h.userService.Delete(userId)
	if err != nil {
		h.logger.Error(err.Error())
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
