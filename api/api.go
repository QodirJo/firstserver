package api

import (
	_ "github.com/Exam4/ApiGatawey/api/docs"
	"github.com/Exam4/ApiGatawey/api/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ApiGatawey
// @description  ApiGatawey
// @version 1.0
// @host localhost:55555
// @BasePath /
// @in header
func NewGin(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	genetic := router.Group("/genetic")
	{
		genetic.POST("/create", h.CreateGenetic)
		genetic.PUT("/update/:id", h.UpdateGenetic) // Updated route
		genetic.GET("/get/:id", h.GetGenetic)
		genetic.DELETE("/delete/:id", h.DeleteGenetic)
		genetic.GET("/list", h.ListGenetic)
	}

	records := router.Group("/records")
	{
		records.POST("/create", h.CreateRecords)
		records.PUT("/update/:id", h.UpdateRecords) // Updated route
		records.GET("/get/:id", h.GetRecords)
		records.DELETE("/delete/:id", h.DeleteRecords)
		records.GET("/list", h.ListRecords)
	}

	lifestyle := router.Group("/lifestyle")
	{
		lifestyle.POST("/create", h.CreateLifeStyle)
		lifestyle.PUT("/update/:id", h.UpdateLifeStyle) // Updated route
		lifestyle.GET("/get/:id", h.GetLifeStyle)
		lifestyle.DELETE("/delete/:id", h.DeleteLifeStyle)
		lifestyle.GET("/list", h.ListLifeStyle)
	}

	wearable := router.Group("/wearable")
	{
		wearable.POST("/create", h.CreateWearable)
		wearable.PUT("/update/:id", h.UpdateWearable) // Updated route
		wearable.GET("/get/:id", h.GetWearable)
		wearable.DELETE("/delete/:id", h.DeleteWearable)
		wearable.GET("/list", h.ListWearable)
	}

	healthRecommendations := router.Group("/health-recommendation")
	{
		healthRecommendations.POST("/create", h.CreateHealthRecommendation)
		healthRecommendations.PUT("/update/:id", h.UpdateHealthRecommendation) // Updated route
		healthRecommendations.GET("/get/:id", h.GetHealthRecommendation)
		healthRecommendations.DELETE("/delete/:id", h.DeleteHealthRecommendation)
		healthRecommendations.GET("/list", h.ListHealthRecommendation)
	}

	return router
}
