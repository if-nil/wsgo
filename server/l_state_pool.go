/*
Copyright © 2023 ifNil ifnil.git@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	l := lua.NewState()
	if err := l.DoFile(pl.luaFile); err != nil {
		logger.Panicf("load lua file failed: %v", err)
	}
	m := l.Get(-1)
	l.SetGlobal("M", m)
	return l
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
