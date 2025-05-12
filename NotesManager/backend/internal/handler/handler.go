package handler

import (
	"github.com/berezovskyivalerii/notes-manager/backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/gin-contrib/cors"
	
	_ "github.com/berezovskyivalerii/notes-manager/backend/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.RedirectTrailingSlash = false
	router.RemoveExtraSlash = true

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("", h.getAllLists)
			lists.POST("", h.createList)

			lists.GET("/:listId", h.getById)
			lists.PUT("/:listId", h.updateList)
			lists.DELETE("/:listId", h.deleteList)

			lists.GET("/:listId/items", h.getAllItems)
			lists.POST("/:listId/items", h.createItem)
		}

		items := api.Group("/items")
		{
			items.GET("/:itemId", h.getItemById)
			items.PUT("/:itemId", h.updateItem)
			items.DELETE("/:itemId", h.deleteItem)
		}
	}

	return router
}
