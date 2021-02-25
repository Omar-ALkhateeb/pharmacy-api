package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func ProductImport(c *gin.Context) {
	db := database.DBConn
	var productsList []types.Product

	f, err := c.FormFile("F1") // the same as getting the parameters carried from the request
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		importFile, err := os.Open(f.Filename)
		defer importFile.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if err := gocsv.UnmarshalFile(importFile, &productsList); err != nil { // Load products from file
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		fmt.Println(&productsList)
		responseMessage := ""
		if err := db.Create(&productsList).Error; err != nil {
			responseMessage = responseMessage + ", " + err.Error()
			c.String(422, responseMessage)
		} else {
			c.String(201, "created")
		}

	}

}
