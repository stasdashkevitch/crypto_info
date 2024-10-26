package websocket

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	cryptotrackerusecase "github.com/stasdashkevitch/crypto_info/internal/usecase/cryptoTrackerUsecase"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type cryptoTrackerWebSocketHandler struct {
	usecase *cryptotrackerusecase.CryptoTrackerUsecase
	logger  *zap.SugaredLogger
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func NewCryptoTrackerWebsocketHandler(handler *http.ServeMux, logger *zap.SugaredLogger, usecase cryptotrackerusecase.CryptoTrackerUsecase) {
	wh := &cryptoTrackerWebSocketHandler{
		usecase: &usecase,
		logger:  logger,
		clients: make(map[*websocket.Conn]bool),
	}

	handler.HandleFunc("GET /ws/crypto/price/{id}", wh.HandleConnections)
}

func (wh *cryptoTrackerWebSocketHandler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	defer func() {
		ws.Close()
	}()
	wh.mu.Lock()
	wh.clients[ws] = true
	wh.mu.Unlock()

	updates, err := wh.usecase.SubscribeCryptoDataPrice(ctx)
	if err != nil {
		wh.logger.Errorf("Redis subscription error: %v", err)
		return
	}

	go func() {
		for {
			wh.logger.Info("start BroadcastUpdate")
			select {
			case <-ctx.Done():
				wh.logger.Info("Update finish")
				return
			case data := <-updates:
				wh.BroadcastUpdate(data)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				wh.logger.Info("Finish per update data")
				return
			default:
				wh.logger.Info("update data")
				err := wh.usecase.UpdateCryptoDataPrice(ctx, id)

				if err != nil {
					wh.logger.Errorf("Error data updating: %v", err)
				}

				wh.logger.Info("ok")
				time.Sleep(10 * time.Second)
			}
		}
	}()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			wh.logger.Infof("Клиент отключился: %v", err)
			break
		}
	}

	defer func() {
		wh.mu.Lock()
		delete(wh.clients, ws)
		wh.mu.Unlock()
		wh.logger.Info("delete client")
	}()
}

func (wh *cryptoTrackerWebSocketHandler) BroadcastUpdate(data []byte) {
	wh.logger.Info(len(wh.clients))
	wh.mu.Lock()
	defer wh.mu.Unlock()

	for client := range wh.clients {
		wh.logger.Info("start write")
		err := client.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			wh.logger.Errorf("Error while sending data: %v", err)
			client.Close()
			delete(wh.clients, client)
			return
		} else {
			wh.logger.Info("data sent to client")
		}
	}
}
