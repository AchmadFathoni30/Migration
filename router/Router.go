package router

import (
	"Migration/pajakDispenda/controller"
	"Migration/salesBandung/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/pajak", controller.MigrateData)
	r.GET("/salesHeader", controllers.GetSalesHeader)
	r.GET("/salesVoid", controllers.GetSalesVoid)
	r.GET("/salesPayment", controllers.GetSalesPayment)
	r.GET("/salesDetail", controllers.GetSalesDetail)
	return r
}
