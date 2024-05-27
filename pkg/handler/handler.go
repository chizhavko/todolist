package handler

import (
	"github.com/chizhavko/todolist/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *service.Services
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	api := router.Group("/api", h.identifyUser)
	{
		lists := api.Group("lists")
		{
			lists.GET("/", h.getAllLists)
			lists.POST("/", h.createList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		items := lists.Group(":id/items")
		{
			items.GET("/", h.getAllItems)
			items.POST("/", h.createItem)
		}
	}

	items := router.Group("/items", h.identifyUser)
	{
		items.GET("/:id", h.getItemById)
		items.PUT("/:id", h.updateItem)
		items.DELETE("/:id", h.deleteItem)
	}

	return router
}
