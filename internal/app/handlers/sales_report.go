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

func SalesReport(c *gin.Context) {
	startDate, endDate := services.DecideDate(c)
	salesReports := services.SalesReportCalculate(startDate, endDate)
	bytesBuffer := &bytes.Buffer{}
	csvWriter := csv.NewWriter(bytesBuffer)

	// CSV header
	headerRow := []string{"Order Note", "Date", "Barcode", "Item name", "Qty",
		"Selling Price per Product", "Total Selling Price", "Purchase Price per Product",
		"Total Purchase Price", "Profit and loss"}
	_ = csvWriter.Write(headerRow)

	// CSV content
	for _, rep := range salesReports {
		row := []string{rep.SalesNote, rep.Time.String(), rep.ProductBarcode,
			rep.ProductName, fmt.Sprint(rep.Quantity), fmt.Sprintf("USD %.2f", rep.SellPricePerProduct),
			fmt.Sprintf("USD %.2f", rep.TotalSellPricePerProduct), fmt.Sprintf("USD %.2f", rep.BuyPricePerProduct),
			fmt.Sprintf("USD %.2f", rep.TotalBuyPricePerProduct), fmt.Sprintf("USD %.2f", rep.ProfitOrLoss)}
		_ = csvWriter.Write(row)
	}
	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		log.Fatal(err)
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=sales_report.csv")
	c.Data(http.StatusOK, "text/csv", bytesBuffer.Bytes())
}

func SalesReportJson(c *gin.Context) {
	startDate, endDate := services.DecideDate(c)
	salesReports := services.SalesReportCalculate(startDate, endDate)

	c.JSON(200, salesReports)
}
