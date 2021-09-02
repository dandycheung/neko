package websocket

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"demodesk/neko/internal/types"
	"demodesk/neko/internal/types/event"
	"demodesk/neko/internal/types/message"
)

type WebSocketPeerCtx struct {
	mu         sync.Mutex
	logger     zerolog.Logger
	connection *websocket.Conn
}

func newPeer(connection *websocket.Conn) *WebSocketPeerCtx {
	logger := log.With().
		Str("module", "websocket").
		Str("submodule", "peer").
		Logger()

	return &WebSocketPeerCtx{
		logger:     logger,
		connection: connection,
	}
}

func (peer *WebSocketPeerCtx) setSessionID(sessionId string) {
	peer.logger = peer.logger.With().Str("session_id", sessionId).Logger()
}

func (peer *WebSocketPeerCtx) Send(event string, payload interface{}) {
	peer.mu.Lock()
	defer peer.mu.Unlock()

	if peer.connection == nil {
		return
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		peer.logger.Err(err).Str("event", event).Msg("message marshalling has failed")
		return
	}

	err = peer.connection.WriteJSON(types.WebSocketMessage{
		Event:   event,
		Payload: raw,
	})

	if err != nil {
		peer.logger.Err(err).Str("event", event).Msg("send message error")
		return
	}

	peer.logger.Debug().
		Str("address", peer.connection.RemoteAddr().String()).
		Str("event", event).
		Str("payload", string(raw)).
		Msg("sending message to client")
}

func (peer *WebSocketPeerCtx) Destroy(reason string) {
	peer.mu.Lock()
	defer peer.mu.Unlock()

	if peer.connection == nil {
		return
	}

	peer.Send(
		event.SYSTEM_DISCONNECT,
		message.SystemDisconnect{
			Message: reason,
		})

	err := peer.connection.Close()
	peer.logger.Err(err).Msg("peer connection destroyed")

	peer.connection = nil
}
