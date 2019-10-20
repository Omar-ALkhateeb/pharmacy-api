package handlers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/dwahyudi/inventory/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InventoryValuation(c *gin.Context) {
	inventoryValuations := services.InventoryValuationCalculate()

	bytesBuffer := &bytes.Buffer{}
	csvWriter := csv.NewWriter(bytesBuffer)

	// CSV header
	headerRow := []string{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"}
	_ = csvWriter.Write(headerRow)

	// CSV Content
	for _, iv := range inventoryValuations {
		row := []string{iv.ProductSku, iv.ProductName, fmt.Sprintf("%d", iv.ProductQuantity),
			fmt.Sprintf("IDR %.2f", iv.ProductAvgPurchasePrice),
			fmt.Sprintf("IDR %.2f", iv.ProductTotalPurchasePrice)}
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
