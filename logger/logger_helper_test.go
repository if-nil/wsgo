package logger

import "testing"

func TestSendLog(t *testing.T) {
	SendLog(TextMessage, []byte("aaaa"))
	RecLog(TextMessage, []byte("aaaa"))
}
