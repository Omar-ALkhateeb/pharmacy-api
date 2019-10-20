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

func SalesReport(c *gin.Context) {
	salesReports := services.SalesReportCalculate()
	bytesBuffer := &bytes.Buffer{}
	csvWriter := csv.NewWriter(bytesBuffer)

	// CSV header
	headerRow := []string{"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah",
		"Harga Jual per Produk", "Total Harga Jual", "Harga Beli per Produk",
		"Total Harga Beli", "Laba/Rugi"}
	_ = csvWriter.Write(headerRow)

	// CSV content
	for _, rep := range salesReports {
		row := []string{rep.SalesNote, rep.Time.String(), rep.ProductSku,
			rep.ProductName, string(rep.Quantity), fmt.Sprintf("%f", rep.SellPricePerProduct),
			fmt.Sprintf("%f", rep.TotalSellPricePerProduct), fmt.Sprintf("%f", rep.BuyPricePerProduct),
			fmt.Sprintf("%f", rep.TotalBuyPricePerProduct), fmt.Sprintf("%f", rep.ProfitOrLoss)}
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
