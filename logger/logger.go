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
	"fmt"
	"os"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	PANIC
	FATAL
	TABWRITER
)

type Logger interface {
	Println(v ...any)
	Printf(format string, v ...any)
}

var (
	defaultLogger Logger
	prefixMap     = map[Level]string{
		DEBUG:   "<cyan>[DEBUG]</>",
		INFO:    "<green>[INFO]</>",
		WARNING: "<yellow>[WARN]</>",
		ERROR:   "<red>[ERROR]</>",
		PANIC:   "<red>[PANIC]</>",
		FATAL:   "<red>[PANIC]</>",
	}
)

func init() {
	defaultLogger = &colorPrinter{}
}

func Log(level Level, v ...any) {
	switch level {
	case PANIC:
		defaultLogger.Printf("%s %s", time.Now().Format("15:04:05.000"), prefixMap[level])
		s := fmt.Sprintln(v...)
		defaultLogger.Println(s)
		panic(s)
	case FATAL:
		defaultLogger.Printf("%s %s", time.Now().Format("15:04:05.000"), prefixMap[level])
		defaultLogger.Println(v...)
		os.Exit(1)
	case TABWRITER:
		defaultLogger.Println(v...)
	// case DEBUG, INFO, WARNING, ERROR, TABWRITER:
	default:
		defaultLogger.Printf("%s %s", time.Now().Format("15:04:05.000"), prefixMap[level])
		defaultLogger.Println(v...)
	}
}

func Logf(level Level, format string, v ...any) {
	format = format + "\n"
	switch level {
	case PANIC:
		defaultLogger.Printf("%s %s", time.Now().Format("15:04:05.000"), prefixMap[level])
		s := fmt.Sprintf(format, v...)
		defaultLogger.Printf("[PANIC]")
		defaultLogger.Printf(s)
		panic(s)
	case FATAL:
		defaultLogger.Printf("%s %s", time.Now().Format("15:04:05.000"), prefixMap[level])
		defaultLogger.Printf(format, v...)
		os.Exit(1)
	// case DEBUG, INFO, WARNING, ERROR, TABWRITER:
	case TABWRITER:
		defaultLogger.Printf(format, v...)
	default:
		defaultLogger.Printf("%s %s", time.Now().Format("15:04:05.000"), prefixMap[level])
		defaultLogger.Printf(format, v...)
	}
}
