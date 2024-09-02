package controllers

import "github.com/gin-gonic/gin"




func GetInvoices() gin.HandlerFunc  {
	return func(c *gin.Context) {
        // get all invoices from db
    }
	
}
func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}


func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

    }
}



func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

    }

}

func DeleteInvoice() gin.HandlerFunc {
return func(c *gin.Context){
	
}
}
// route.GET("/invoices", controller.GetInvoices())
// 	route.POST("/invoices", controller.CreateInvoice())
// 	route.GET("/invoices/:id", controller.GetInvoice())
// 	route.PATCH("/invoices/:id", controller.UpdateInvoice())
// 	route.DELETE("/invoices/:id", controller.DeleteInvoice())