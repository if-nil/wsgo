package server

import (
	"github.com/gorilla/websocket"
	"github.com/if-nil/wsgo/logger"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
}

type ServerType int

const (
	Echo ServerType = iota
)

type Server struct {
	ServerType ServerType
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch s.ServerType {
	case Echo:
		s.echo(w, r)
	}
}

func (s *Server) echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	defer c.Close()
	c.SetPingHandler(func(data string) error {
		logger.RecLog(logger.PingMessage, []byte(data))
		if err := c.WriteControl(websocket.PongMessage, []byte(data), time.Now().Add(time.Second*10)); err != nil {
			logger.Error("Write Pong failed", err)
			c.Close()
		}
		logger.SendLog(logger.PongMessage, []byte(data))
		return nil
	})
	c.SetPongHandler(func(data string) error {
		logger.RecLog(logger.PongMessage, []byte(data))
		return nil
	})
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			logger.Error(err)
			break
		}
		logger.RecLog(mt, message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			logger.Error(err)
			break
		}
		logger.SendLog(mt, message)
	}
}
