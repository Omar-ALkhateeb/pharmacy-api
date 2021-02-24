package handlers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	"github.com/Omar-ALkhateeb/pharm-inventory/internal/services"
	"github.com/gin-gonic/gin"
)

func InventoryValuation(c *gin.Context) {
	inventoryValuations := services.InventoryValuationCalculate()

	bytesBuffer := &bytes.Buffer{}
	csvWriter := csv.NewWriter(bytesBuffer)

	// CSV header
	headerRow := []string{"Barcode", "Item Name", "Qty", "Avg selling price", "Total"}
	_ = csvWriter.Write(headerRow)

	// CSV Content
	for _, iv := range inventoryValuations {
		row := []string{iv.ProductBarcode, iv.ProductName, fmt.Sprintf("%d", iv.ProductQuantity),
			fmt.Sprintf("USD %.2f", iv.ProductAvgPurchasePrice),
			fmt.Sprintf("USD %.2f", iv.ProductTotalPurchasePrice)}
		_ = csvWriter.Write(row)
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		log.Fatal(err)
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=inventory_valuation.csv")
	c.Data(http.StatusOK, "text/csv", bytesBuffer.Bytes())
}
