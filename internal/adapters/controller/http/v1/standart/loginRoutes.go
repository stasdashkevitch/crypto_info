package standart

import (
	"encoding/json"
	"net/http"

	"github.com/stasdashkevitch/crypto_info/internal/dtos"
	"github.com/stasdashkevitch/crypto_info/internal/usecase"
	"go.uber.org/zap"
)

type loginHandler struct {
	usecase usecase.LoginServis
	logger  *zap.SugaredLogger
}

func NewLoginHandler(handler *http.ServeMux, logger *zap.SugaredLogger, useusecase usecase.LoginServis) {
	h := &loginHandler{
		usecase: useusecase,
		logger:  logger,
	}
	handler.HandleFunc("POST /login", h.Login)
}

func (h *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("login")
	var req dtos.LoginUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	resp, err := h.usecase.Login(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Location", "/")
	json.NewEncoder(w).Encode(resp)
}
