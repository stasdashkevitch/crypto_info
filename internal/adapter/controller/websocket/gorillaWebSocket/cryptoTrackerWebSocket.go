package gorillawebsocket

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
	wh.logger.Infow("Request for websocket connection: ",
		"url", r.URL)
	id := r.PathValue("id")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ws, err := upgrader.Upgrade(w, r, nil)
	wh.logger.Info("Upgrade to websocket")
	if err != nil {
		wh.logger.Errorf("Error updating to websocket: %v", err)
		return
	}

	defer func() {
		wh.logger.Info("Closing a websocket connection")
		ws.Close()
	}()
	wh.mu.Lock()
	wh.logger.Info("Add websocket client")
	wh.clients[ws] = true
	wh.mu.Unlock()

	wh.logger.Info("Redis subscription to crypto data price")
	updates, err := wh.usecase.SubscribeCryptoDataPrice(ctx)
	if err != nil {
		wh.logger.Errorf("Redis subscription error: %v", err)
		return
	}

	go func() {
		wh.logger.Info("Start BroadcastUpdate")
		for {
			select {
			case <-ctx.Done():
				wh.logger.Info("BroadcastUpdate finish")
				return
			case data := <-updates:
				wh.logger.Info("Recieve some data for BroadcastUpdate")
				wh.BroadcastUpdate(data)
			}
		}
	}()

	go func() {
		wh.logger.Info("Start UpdateCryptoDataPrice")
		for {
			select {
			case <-ctx.Done():
				wh.logger.Info("Finish per update data")
				return
			default:
				wh.logger.Info("Update data")
				err := wh.usecase.UpdateCryptoDataPrice(ctx, id)
				if err != nil {
					wh.logger.Errorf("Error data updating: %v", err)
				}

				wh.logger.Info("Timeoute 1 minute")
				time.Sleep(1 * time.Minute)
			}
		}
	}()

	for {
		wh.logger.Info("Read message from client")
		_, _, err := ws.ReadMessage()
		if err != nil {
			wh.logger.Infof("The client is disconnected: %v", err)
			break
		}
	}

	defer func() {
		wh.mu.Lock()
		wh.logger.Info("Deleting a client")
		delete(wh.clients, ws)
		wh.mu.Unlock()
	}()
}

func (wh *cryptoTrackerWebSocketHandler) BroadcastUpdate(data []byte) {
	wh.logger.Info("Current number of clients:", len(wh.clients))
	wh.mu.Lock()
	defer wh.mu.Unlock()

	wh.logger.Info("Start send data to client")
	for client := range wh.clients {
		err := client.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			wh.logger.Errorf("Error while sending data: %v", err)
			wh.logger.Info("Closing a websocket connection")
			client.Close()
			wh.logger.Info("Deleting a client")
			delete(wh.clients, client)
			return
		} else {
			wh.logger.Info("Data sent to client")
		}
	}
}
