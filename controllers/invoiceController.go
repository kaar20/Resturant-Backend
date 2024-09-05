package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaar20/resturant_backend/database"
	"github.com/kaar20/resturant_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Payment_due_date time.Time
	Order_details    interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoices")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		var allInvoices []bson.M

		if err := result.All(ctx, &allInvoices); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, allInvoices)

		// get all invoices from db
	}

}
func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        var invoice models.Invoice
        
        if err := c.BindJSON(&invoice); err!= nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
		var order models.Order

		err := ordersCollection.FindOne(ctx, bson.M{"order_id": invoice.Order_id}).Decode(&order)
		defer cancel()
		if err!= nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
            return
        }
	status := "PENDING"
	if invoice.Payment_status != nil{
		invoice.Payment_status = &status	
	}

	invoice.Payment_due_date, _ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	invoice.Created_at, _ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	invoice.Updated_at, _ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	invoice.ID = primitive.NewObjectID()
	invoice.Invoice_id = invoice.ID.Hex()

	validation := validate.Struct(invoice)
	if validation !=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Invoice Error: ":validation})
		return
	}
	 result , err := invoiceCollection.InsertOne(ctx,invoice)
	 if err!= nil {
         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
         return
     }
	 defer cancel()
	 c.JSON(http.StatusOK,result)




        // result, err := invoiceCollection.InsertOne(ctx, invoice)
        // if err!= nil {
        //     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        //     return
        // }

        // invoice.Id = result.InsertedID.(primitive.ObjectID).Hex()
        // c.JSON(http.StatusOK, invoice)

        // insert invoice into db

	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var invoiceId = c.Param("id")
		var invoice models.Invoice
		err := invoiceCollection.FindOne(ctx, bson.M{"id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// c.JSON(http.StatusOK, invoice)
		var invoiceView InvoiceViewFormat

		allOrderItems, err := ItemsByOrder(invoice.Order_id)
		invoiceView.Order_id = invoice.Order_id
		invoiceView.Payment_due_date = invoice.Payment_due_date
		invoiceView.Payment_method = "null"

		if invoiceView.Payment_method != "" {
			invoiceView.Payment_method = *invoice.Payment_method
		}
		invoiceView.Invoice_id = invoice.Invoice_id
		invoiceView.Payment_status = *&invoice.Payment_status
		invoiceView.Payment_due = allOrderItems[0]["payment_due"]
		invoiceView.Table_number = allOrderItems[0]["table_number"]
		invoiceView.Order_details = allOrderItems[0]["order_details"]

		c.JSON(http.StatusOK, invoiceView)

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		var invoiceId = c.Param("id")
		var invoice models.Invoice
		var updateObj primitive.D

		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		filter := bson.M{"invoice_id": invoiceId}
		if invoice.Payment_method != nil {
			updateObj = append(updateObj, bson.E{"payment_method", invoice.Payment_method})

		}
		if invoice.Payment_status != nil {
			updateObj = append(updateObj, bson.E{"payment_status", invoice.Payment_status})

		}
		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", invoice.Updated_at})
		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		status := "PENDING"
		if invoice.Payment_status != nil {
			invoice.Payment_status = &status

		}

		result, err := invoiceCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}},
			&opt)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}

}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// route.GET("/invoices", controller.GetInvoices())
// 	route.POST("/invoices", controller.CreateInvoice())
// 	route.GET("/invoices/:id", controller.GetInvoice())
// 	route.PATCH("/invoices/:id", controller.UpdateInvoice())
// 	route.DELETE("/invoices/:id", controller.DeleteInvoice())
