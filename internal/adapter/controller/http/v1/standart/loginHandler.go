package standart

import (
	"encoding/json"
	"net/http"

	"github.com/stasdashkevitch/crypto_info/internal/dto"
	loginusecase "github.com/stasdashkevitch/crypto_info/internal/usecase/loginUsecase"
	"go.uber.org/zap"
)

type loginHandler struct {
	usecase *loginusecase.LoginUsecase
	logger  *zap.SugaredLogger
}

func NewLoginHandler(handler *http.ServeMux, logger *zap.SugaredLogger, usecase *loginusecase.LoginUsecase) {
	h := &loginHandler{
		usecase: usecase,
		logger:  logger,
	}
	handler.HandleFunc("POST /login", h.Login)
}

func (h *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("Recieved request: ",
		"method", r.Method,
		"url", r.URL)
	var req dto.LoginUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Invalid input: ", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	resp, err := h.usecase.Login(req)

	if err != nil {
		h.logger.Error("Failed login attempt: ", err)
		http.Error(w, "Failed to login", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Location", "/")
	json.NewEncoder(w).Encode(resp)
}
