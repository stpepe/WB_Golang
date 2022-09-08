package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stpepe/nats-task/pkg/service"
	"net/http"
)

type Handler struct{
	services *service.Service
	publisher *service.PubNATS
	cache *service.Cache
	channel chan int
}

func NewHandler(services *service.Service, publisher *service.PubNATS, cache *service.Cache, channel chan int) *Handler{
	n := Handler{
		services: services,
		publisher: publisher,
		cache: cache,
		channel: channel,
	}
	return &n
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()
	router.LoadHTMLGlob("static/*")
	orders := router.Group("/order")
	{
		orders.GET("/:id", h.show)
		orders.POST("/send", h.send)
		orders.GET("/cache", h.rec_cache)
		orders.GET("/", func(c *gin.Context){
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}
	return router
}