package cron

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/robfig/cron/v3"
)

type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	// Error logs an error condition.
	Error(err error, msg string, keysAndValues ...interface{})
}

var C *cron.Cron

type L struct {
	out io.Writer
}

func NewL() *L {
	return &L{
		out: os.Stdout,
	}
}

func (l *L) Info(msg string, keysAndValues ...interface{}) {
	bf := bytes.Buffer{}
	for i := 0; i < len(keysAndValues)/2; i++ {
		if i > 0 {
			bf.WriteString(", ")
		}
		bf.WriteString("%v=%v")
	}
	l.out.Write([]byte(fmt.Sprintf("my logger %s %s %s", msg, bf.String(), keysAndValues)))

}

func (l *L) Error(err error, msg string, keysAndValues ...interface{}) {
	bf := bytes.Buffer{}
	for i := 0; i < len(keysAndValues)/2; i++ {
		if i > 0 {
			bf.WriteString(", ")
		}
		bf.WriteString("%v=%v")
	}
	l.out.Write([]byte(fmt.Sprintf("my logger %s %s %s %s", err, msg, bf.String(), keysAndValues)))
}

func Setup() {
	C = cron.New(cron.WithLogger(cron.DefaultLogger), cron.WithSeconds())

	C.AddFunc("* * * * * *", func() {
		fmt.Println("wawawa")
	})
	C.Start()
}
