package logs

import (
	"io"
	"log/slog"
	"os"
	"time"
)

var (
	AL *slog.Logger
	SL *slog.Logger
)

var opts = slog.HandlerOptions{
	AddSource: true,
	Level:     slog.LevelDebug,
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Value.Kind() == slog.KindTime {
			return slog.String(a.Key, a.Value.Time().Format(time.RFC3339))
		}
		return a
	},
}

func SetupLogger() {
	f1, err := os.Create("./logs/access.log")
	if err != nil {
		panic(err)
	}
	f2, err := os.Create("./logs/app.log")
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	f1.Close()
	// 	f2.Close()
	// }()
	h1 := slog.NewJSONHandler((io.MultiWriter(f1, os.Stdout)), &opts)
	h2 := slog.NewTextHandler((io.MultiWriter(f2, os.Stdout)), &opts)
	
	AL = slog.New(h1).WithGroup("access1")
	SL = slog.New(h2).WithGroup("service1")

}
