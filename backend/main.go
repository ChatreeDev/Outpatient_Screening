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
			protected.PATCH("/Nurses", controller.UpdateNurse)
			protected.DELETE("/Nurses/:id", controller.DeleteNurse)

			// HistorySheet Routes
			protected.GET("/HistorySheets", controller.ListHistorySheets)
			protected.GET("/HistorySheet/:id", controller.GetHistorySheet)

			// DiabetesLevel Routes
			protected.GET("/DiabetesLevels", controller.ListDiabetesLevels)
			protected.GET("/DiabetesLevel/:id", controller.GetDiabetesLevel)
			protected.POST("/DiabetesLevels", controller.CreateDiabetesLevel)
			protected.PATCH("/DiabetesLevels", controller.UpdateDiabetesLevel)
			protected.DELETE("/DiabetesLevels/:id", controller.DeleteDiabetesLevel)

			// HighBloodPressureLevel Routes
			protected.GET("/highbloodpressure_levels", controller.ListHighBloodPressureLevels)
			protected.GET("/highbloodpressure_level/:id", controller.GetHighBloodPressureLevel)
			protected.POST("/highbloodpressure_levels", controller.CreateHighBloodPressureLevel)
			protected.PATCH("/highbloodpressure_levels", controller.UpdateHighBloodPressureLevel)
			protected.DELETE("/highbloodpressure_levels/:id", controller.DeleteHighBloodPressureLevel)

			// OutpatientScreening Routes
			protected.GET("/OutpatientScreenings", controller.ListOutpatientScreenings)
			protected.GET("/OutpatientScreening/:id", controller.GetOutpatientScreening)
			protected.POST("/OutpatientScreenings", controller.CreateOutpatientScreening)
			// protected.PATCH("/Nurses", controller.UpdateNurse)
			// protected.DELETE("/Nurses/:id", controller.DeleteNurse
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
