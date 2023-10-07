package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_ServeHTTP(t *testing.T) {
	s := &Server{
		ServerType: LuaServer,
		LuaPool:    New("../example/upgrade.lua"),
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	s.ServeHTTP(w, r)
	resp := w.Result()
	t.Logf("%#v", resp)
}
