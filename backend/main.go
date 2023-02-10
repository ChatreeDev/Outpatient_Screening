package main

import (
	"github.com/ChatreeDev/sa-65-example/controller"
	"github.com/ChatreeDev/sa-65-example/middlewares"

	"github.com/ChatreeDev/sa-65-example/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("/")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Nurse Routes
			protected.GET("/Nurses", controller.ListNurses)
			protected.GET("/Nurse/:id", controller.GetNurse)
			protected.POST("/Nurses", controller.CreateNurse)

			// HistorySheet Routes
			protected.GET("/HistorySheets", controller.ListHistorySheets)
			protected.GET("/HistorySheet/:id", controller.GetHistorySheet)
			protected.POST("/HistorySheets", controller.CreateHistorySheet)

			// Emergency Routes
			protected.GET("/emergency_levels", controller.ListEmergencyLevels)
			protected.GET("/emergency_level/:id", controller.GetEmergencyLevel)
			protected.POST("/emergency_levels", controller.CreateEmergencyLevel)

			// DiabetesLevel Routes
			protected.GET("/DiabetesLevels", controller.ListDiabetesLevels)
			protected.GET("/DiabetesLevel/:id", controller.GetDiabetesLevel)
			protected.POST("/DiabetesLevels", controller.CreateDiabetesLevel)

			// HighBloodPressureLevel Routes
			protected.GET("/highbloodpressure_levels", controller.ListHighBloodPressureLevels)
			protected.GET("/highbloodpressure_level/:id", controller.GetHighBloodPressureLevel)
			protected.POST("/highbloodpressure_levels", controller.CreateHighBloodPressureLevel)

			// ObesityLevel Routes
			protected.GET("/ObesityLevels", controller.ListObesityLevels)
			protected.GET("/obesity_level/:id", controller.GetObesityLevel)
			protected.POST("/ObesityLevels", controller.CreateObesityLevel)

			// OutpatientScreening Routes
			protected.GET("/outpatientScreenings", controller.ListOutpatientScreenings)
			protected.GET("/outpatientScreening/:id", controller.GetOutpatientScreening)
			protected.POST("/outpatientScreenings", controller.CreateOutpatientScreenings)

		}
	}

	r.POST("/login", controller.Login)

	// Run the server
	r.Run()

}
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
