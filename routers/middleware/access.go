package middleware

import (
	"fmt"
	"gin/logs"
	"log/slog"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type TraceID string

func AccessLog() gin.HandlerFunc {
	rand.NewSource(time.Now().UnixMicro())
	return func(ctx *gin.Context) {
		trace_id := fmt.Sprintf("unique key %d", 10000+rand.Intn(10000))
		ctx.Set("trace_id", trace_id)
		t1 := time.Now()
		ctx.Next()
		latency := time.Since(t1).Milliseconds()
		logs.AL.InfoContext(ctx, "enter middleware", "trace_id", trace_id, slog.String("latency", fmt.Sprintf("%d ms", latency)))
	}

}
