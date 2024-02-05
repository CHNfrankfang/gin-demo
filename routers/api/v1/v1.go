package v1

import (
	"fmt"
	"gin/logs"
	"gin/routers/client"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type HealthCheck struct {
}

type Scan struct {
	SC *client.ScanClient
}

type Cron struct {
	Cron *cron.Cron
}

func (cc *Cron) Acitons(c *gin.Context) {

	action := c.Query("action")

	switch action {
	case "start":
		cc.Start()
		c.JSON(200, "ok")
	case "stop":
		cc.Stop()
		c.JSON(200, "ok")
	case "info":
		scroll := `
	<body style="word-wrap: break-word; white-space: pre-wrap;"><script>
		// https://stackoverflow.com/questions/14866775/detect-document-height-change
		// Purpose: Make sure the scroll bar is always at the bottom of the page as the page continues to output.
		// create an Observer instance
		const resizeObserver = new ResizeObserver(() =>
			window.scrollTo(0, document.body.scrollHeight)
		)
		// start observing a DOM node
		resizeObserver.observe(document.body)
	</script></body>`
		fmt.Fprint(c.Writer, scroll)
		flusher, ok := c.Writer.(http.Flusher)
		if !ok {
			http.Error(c.Writer, "Streaming not supported", http.StatusInternalServerError)
			return
		}
		ctx := c.Request.Context()
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		for {
			select {
			case <-ctx.Done():
				log.Print("Client disconnected")
				return
			default:
				fmt.Fprintf(c.Writer, "data: %s\n", time.Now().Format(time.RFC3339))
				flusher.Flush()

			}
			time.Sleep(time.Second)

		}

	}

}

func (cc *Cron) Start() {
	cc.Cron.Start()
}

func (cc *Cron) Stop() {
	cc.Cron.Stop()
}
func (cc *Cron) Info() map[int]interface{} {
	m := make(map[int]interface{})
	for _, v := range cc.Cron.Entries() {
		m[int(v.ID)] = map[string]interface{}{
			"prev": v.Prev,
			"next": v.Next,
		}
	}
	return m
}

func (hc HealthCheck) Pong(a string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		time.Sleep(1 * time.Second)
		t, _ := ctx.Get("trace_id")
		logs.SL.InfoContext(ctx, "service", "trace_id", t)
		ctx.JSON(200, gin.H{
			"msg":      "ok",
			"data":     "",
			"trace_id": t,
		})
	}
}

func (s Scan) Actions(c *gin.Context) {

	action := c.Query("action")

	switch action {
	case "info":
		c.JSON(200, s.SC)
		return
	case "start":
		c.JSON(200, s.SC.Start("/Users/didi/Documents/goProject/src/gin/logs"))
	case "stop":
		c.JSON(200, s.SC.Stop())
	default:
		c.JSON(500, "params error")
	}

}
