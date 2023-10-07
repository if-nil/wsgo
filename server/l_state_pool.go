package server

import (
	"github.com/if-nil/wsgo/logger"
	"github.com/yuin/gopher-lua"
	"sync"
)

type LStatePool struct {
	m       sync.Mutex
	saved   []*lua.LState
	luaFile string
}

func New(luaFile string) *LStatePool {
	return &LStatePool{
		luaFile: luaFile,
		saved:   make([]*lua.LState, 0, 4),
	}
}

func (pl *LStatePool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.saved)
	if n == 0 {
		return pl.New()
	}
	x := pl.saved[n-1]
	pl.saved = pl.saved[0 : n-1]
	return x
}

func (pl *LStatePool) New() *lua.LState {
	L := lua.NewState()
	if err := L.DoFile(pl.luaFile); err != nil {
		logger.Panicf("load lua file failed: %v", err)
	}
	return L
}

func (pl *LStatePool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.saved = append(pl.saved, L)
}

func (pl *LStatePool) Shutdown() {
	for _, L := range pl.saved {
		L.Close()
	}
}
