package logger

import (
	"testing"
	"time"
)

func TestLogOut(t *testing.T) {
	SetLogger("./log.json")
	Info("🔨 show log info test", "time", time.Now().Unix())
}

func BenchmarkError(b *testing.B) {
	SetLogger("./log.json")
	for i := 0; i < b.N; i++ {
		Info("🔨 show log info test", "time", time.Now().Unix())
	}
}
