package routers

import (
	"gin/cron"
	v1 "gin/routers/api/v1"
	"gin/routers/client"
	"gin/routers/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterHandler() {
	g := gin.Default()
	g.Use(middleware.AccessLog())
	apiV1 := g.Group("api/v1")
	h := v1.HealthCheck{}
	s := v1.Scan{
		SC: &client.ScanClient{},
	}
	c := v1.Cron{
		Cron: cron.C,
	}
	{
		apiV1.GET("ping", h.Pong("a"))
		apiV1.GET("scan/actions", s.Actions)
		apiV1.GET("cron/actions", c.Acitons)
	}

	// go http.ListenAndServe(":8000", g)
	go g.Run(":8001")

}
