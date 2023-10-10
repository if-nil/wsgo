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
package logger

import (
	"time"
)

const (
	TextMessage   = 1
	BinaryMessage = 2
	CloseMessage  = 8
	PingMessage   = 9
	PongMessage   = 10
)

var typMap = map[int]string{
	TextMessage:   "TEXT",
	BinaryMessage: "BINARY",
	CloseMessage:  "CLOSE",
	PingMessage:   "PING",
	PongMessage:   "PONG",
}

func SendLog(typ int, msg []byte) {
	Logf(TABWRITER, "%s <fg=ff5370><<< </><fg=4dd0e1>[%s]</>%s", time.Now().Format("15:04:05.000"), typMap[typ], string(msg))
}

func RecLog(typ int, msg []byte) {
	Logf(TABWRITER, "%s <fg=c3e88d>>>> </><fg=4dd0e1>[%s]</>%s", time.Now().Format("15:04:05.000"), typMap[typ], string(msg))
}

func Debug(v ...any) {
	Log(DEBUG, v...)
}

func Info(v ...any) {
	Log(INFO, v...)
}

func Warning(v ...any) {
	Log(WARNING, v...)
}

func Error(v ...any) {
	Log(ERROR, v...)
}

func Panic(v ...any) {
	Log(PANIC, v...)
}

func Fatal(v ...any) {
	Log(FATAL, v...)
}

func Debugf(format string, v ...any) {
	Logf(DEBUG, format, v...)
}

func Infof(format string, v ...any) {
	Logf(INFO, format, v...)
}

func Warningf(format string, v ...any) {
	Logf(WARNING, format, v...)
}

func Errorf(format string, v ...any) {
	Logf(ERROR, format, v...)
}

func Panicf(format string, v ...any) {
	Logf(PANIC, format, v...)
}

func Fatalf(format string, v ...any) {
	Logf(FATAL, format, v...)
}
