package standart

import (
	"encoding/json"
	"net/http"

	"github.com/stasdashkevitch/crypto_info/internal/dtos"
	"github.com/stasdashkevitch/crypto_info/internal/usecase"
	"go.uber.org/zap"
)

type registrationHandler struct {
	usecase usecase.RegistrationService
	logger  *zap.SugaredLogger
}

func NewRegistrationHandler(handler *http.ServeMux, logger *zap.SugaredLogger, usecase usecase.RegistrationService) {
	h := &registrationHandler{
		usecase: usecase,
		logger:  logger,
	}
	handler.HandleFunc("POST /registration", h.Registration)
}

func (h *registrationHandler) Registration(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("дошло до репы")
	var req dtos.RegisterUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	h.logger.Info("дошло до usecase")

	err := h.usecase.Register(req)

	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Location", "/login")
	w.WriteHeader(http.StatusCreated)
}
