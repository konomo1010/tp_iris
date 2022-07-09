package log

import (
	"bufio"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"os"
	"time"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"tp_iris/config"
)

func MakeAccessLog() *accesslog.AccessLog {
	// Optionally, let's Go with log rotation.
	w, err := rotatelogs.New(
		config.C.Log.File,
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour))
	if err != nil {
		panic(err)
	}

	// Initialize a new access log middleware.
	// Accepts an `io.Writer`.
	ac := accesslog.New(bufio.NewWriter(w))

	ac.Delim = ' '
	ac.Async = false
	ac.BytesReceived = true
	ac.BytesSent = true

	ac.AddOutput(os.Stdout)
	return ac
}
