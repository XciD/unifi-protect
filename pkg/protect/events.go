package protect

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func init() {
	websocket.DefaultDialer.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
}

type WebsocketEvent struct {
	nvr          *NVR
	Events       chan *WsMessage
	socket       *websocket.Conn
	disconnected chan bool
}

func NewWebsocketEvent(nvr *NVR) (*WebsocketEvent, error) {
	unifiProtectWebsocket := &WebsocketEvent{
		nvr:          nvr,
		Events:       make(chan *WsMessage),
		disconnected: make(chan bool),
	}

	if !nvr.connected {
		return nil, errors.New("not connected")
	}

	if err := unifiProtectWebsocket.connect(); err != nil {
		return nil, err
	}

	return unifiProtectWebsocket, nil
}

func (l *WebsocketEvent) connect() error {
	if err := l.connectWs(); err != nil {
		return err
	}

	go l.handleReconnect()

	return nil
}

func (l *WebsocketEvent) connectWs() error {
	log.Info("Connecting to WS")
	u := url.URL{
		Scheme: "wss",
		Host:   fmt.Sprintf("%s:%d", l.nvr.host, l.nvr.port),
		Path:   "/proxy/protect/ws/updates",
	}

	headers := http.Header{}
	headers.Add("Cookie", l.nvr.cookies)
	headers.Add("X-CSRF-Token", l.nvr.csrfToken)

	socket, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		return err
	}

	l.socket = socket

	go l.readPump()

	return nil
}

func (l *WebsocketEvent) readPump() {
	defer func() {
		log.Info("Stopping websocket pump")
		_ = l.socket.Close()
	}()
	log.Info("Starting websocket pump")

	for {
		_, rawMessage, err := l.socket.ReadMessage()
		if err != nil {
			l.disconnected <- true
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message, err := DecodeWsMessage(rawMessage)

		if err != nil {
			log.Errorf("Invalid rawMessage %s", err)
			continue
		}

		log.Trace("Pushing new rawMessage from socket to socket channel")
		l.Events <- message
	}
}

func (l *WebsocketEvent) handleReconnect() {
	// If we finish, we restart a reconnect loop
	defer func() {
		log.Info("Stopping disconnect loop")
	}()

	for {
		select {
		case <-l.disconnected:
			for {
				log.Warn("Disconnected, reconnecting in 30s")
				time.Sleep(30 * time.Second)

				if err := l.connect(); err != nil {
					log.Warnf("Error during reconnection, retrying (%s)", err.Error())
					continue
				}
				break
			}
			break
		}
	}
}
