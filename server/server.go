package server

import (
	"github.com/gorilla/websocket"
	"github.com/if-nil/wsgo/logger"
	lua "github.com/yuin/gopher-lua"
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
	LuaServer
)

type Server struct {
	ServerType ServerType
	LuaPool    *LStatePool
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch s.ServerType {
	case Echo:
		s.echo(w, r)
	case LuaServer:
		s.luaServer(w, r)
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

func (s *Server) luaServer(w http.ResponseWriter, r *http.Request) {
	// upgrade_callback
	// handler
	l := s.LuaPool.Get()
	defer s.LuaPool.Put(l)

	var header http.Header = nil

	upgradeFn := l.GetGlobal("upgrade_callback")
	if upgradeFn.Type() != lua.LTNil {
		if err := l.CallByParam(lua.P{
			Fn:      upgradeFn,
			NRet:    1,
			Protect: true,
		}, &lua.LTable{}); err != nil {
			logger.Error(err)
			return
		}
		v := l.Get(-1) // returned value
		l.Pop(1)       // remove received value
		if v, ok := v.(lua.LBool); !ok || !bool(v) {
			logger.Infof("could not upgrade to a websocket connection")
			return
		}
	}
	c, err := upgrader.Upgrade(w, r, header)
	if err != nil {
		logger.Error(err)
		return
	}
	defer c.Close()
}
