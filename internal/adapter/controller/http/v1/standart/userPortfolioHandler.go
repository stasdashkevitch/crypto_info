package standart

import (
	"encoding/json"
	"net/http"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
	userportfoliousecase "github.com/stasdashkevitch/crypto_info/internal/usecase/userPortfolioUsecase"
	"go.uber.org/zap"
)

type userPortfolioHandler struct {
	usecase *userportfoliousecase.UserPortfolioUsecase
	logger  *zap.SugaredLogger
}

type response struct {
	response map[string]string
}

func NewUserPortfolioHandler(handler *http.ServeMux, logger *zap.SugaredLogger, usecase *userportfoliousecase.UserPortfolioUsecase) {
	h := &userPortfolioHandler{
		usecase: usecase,
		logger:  logger,
	}

	handler.HandleFunc("POST /api/portfolio", h.CreateUserPortfolio)
	handler.HandleFunc("GET /api/portfolio/{user_id}", h.GetAllUserPortfolio)
	handler.HandleFunc("GET /api/portfolio/{user_id}/{crypto_id}", h.GetUserPortfolioByCryptoID)
	handler.HandleFunc("PUT /api/portfolio", h.UpdateUserPortfolio)
	handler.HandleFunc("DELETE /api/portfolio/{user_id}/{crypto_id}", h.DeleteUserPortfolio)
}

func (h *userPortfolioHandler) CreateUserPortfolio(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("Recieved request: ",
		"method", r.Method,
		"url", r.URL)

	var userPortfolio entity.UserPortfolio
	if err := json.NewDecoder(r.Body).Decode(&userPortfolio); err != nil {
		h.logger.Errorf("Error decoding request body: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err := h.usecase.CreateUserPortfolio(&userPortfolio)
	if err != nil {
		h.logger.Errorf("Error creating user portfolio: %v", err)
		http.Error(w, "Error creating user portfolio", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response{map[string]string{"message": "User portfolio created succesfully"}})
}

func (h *userPortfolioHandler) GetAllUserPortfolio(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("Recieved request: ",
		"method", r.Method,
		"url", r.URL)

	userID := r.PathValue("user_id")

	allUserPortfolio, err := h.usecase.GetAllUserPortfolio(userID)
	if err != nil {
		h.logger.Errorf("Invalid input: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(allUserPortfolio); err != nil {
		h.logger.Errorf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *userPortfolioHandler) GetUserPortfolioByCryptoID(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("Recieved request: ",
		"method", r.Method,
		"url", r.URL)

	userID := r.PathValue("user_id")
	cryptoID := r.PathValue("crypto_id")

	userPortfolio, err := h.usecase.GetUserPortfolio(userID, cryptoID)
	if err != nil {
		h.logger.Errorf("Invalid input: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userPortfolio); err != nil {
		h.logger.Errorf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *userPortfolioHandler) UpdateUserPortfolio(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("Recieved request: ",
		"method", r.Method,
		"url", r.URL)

	var userPortfolio entity.UserPortfolio

	if err := json.NewDecoder(r.Body).Decode(&userPortfolio); err != nil {
		h.logger.Errorf("Error decoding request body: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.usecase.UpdateUserPortfolio(&userPortfolio)
	if err != nil {
		h.logger.Errorf("Error updating user portfolio: %v", err)
		http.Error(w, "Error updating user portfolio", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{map[string]string{"message": "User portfolio updated succesfully"}})
}

func (h *userPortfolioHandler) DeleteUserPortfolio(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("Recieved request: ",
		"method", r.Method,
		"url", r.URL)

	userID := r.PathValue("user_id")
	cryptoID := r.PathValue("crypto_id")

	err := h.usecase.DeleteUserPortfolio(userID, cryptoID)
	if err != nil {
		h.logger.Errorf("Error deleting user portfolio: %v", err)
		http.Error(w, "Error deleting user portfolio", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response{map[string]string{"message": "User portfolio deleting succesfully"}})
}
