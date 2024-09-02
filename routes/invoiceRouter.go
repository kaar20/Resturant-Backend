package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/kaar20/resturant_backend/controllers"
)

// import "github.com/kaar20/"

func InvoiceRoutes(route *gin.Engine) {
	route.GET("/invoices", controller.GetInvoices())
	route.POST("/invoices", controller.CreateInvoice())
	route.GET("/invoices/:id", controller.GetInvoice())
	route.PATCH("/invoices/:id", controller.UpdateInvoice())
	route.DELETE("/invoices/:id", controller.DeleteInvoice())
}
