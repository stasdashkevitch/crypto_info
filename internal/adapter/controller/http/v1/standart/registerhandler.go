package standart

import (
	"encoding/json"
	"net/http"

	"github.com/stasdashkevitch/crypto_info/internal/dto"
	registrationusecase "github.com/stasdashkevitch/crypto_info/internal/usecase/registrationUsecase"
	"go.uber.org/zap"
)

type registrationHandler struct {
	usecase registrationusecase.RegistrationUsecase
	logger  *zap.SugaredLogger
}

func NewRegistrationHandler(handler *http.ServeMux, logger *zap.SugaredLogger, usecase registrationusecase.RegistrationUsecase) {
	h := &registrationHandler{
		usecase: usecase,
		logger:  logger,
	}
	handler.HandleFunc("POST /registration", h.Registration)
}

func (h *registrationHandler) Registration(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("Recieved request: ",
		"method", r.Method,
		"url", r.URL)
	var req dto.RegisterUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid input: ", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err := h.usecase.Register(req)

	if err != nil {
		h.logger.Error("Failed to registr: ", err)
		http.Error(w, "Failed to registr", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Location", "/login")
	w.WriteHeader(http.StatusCreated)
}
