package app

import (
	"fmt"
	"log"

	"github.com/certified-juniors/AtomHackFinalEmailService/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Run запускает приложение
func (app *Application) Run() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "AtomHackMarsBackend RestAPI"
	docs.SwaggerInfo.Description = "API server for Mars application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/api/v1"

	// r.Use(middleware.CorsMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	ApiGroup := r.Group("/api/v1")
	{
		DocumentGroup := ApiGroup.Group("/document")
		{
			DocumentGroup.POST("/send-to-support", app.handler.SendToSupport)

		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf("%s:%d", app.cfg.API.ServiceHost, app.cfg.API.ServicePort)
	r.Run(addr)

	log.Println("Server down")
}
