package handlers

import (
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/reporttypes"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/services"
	"github.com/gin-gonic/gin"
)

func SalesReportSummary(c *gin.Context) {
	startDate, endDate := services.DecideDate(c)
	salesReportData := services.SalesReportCalculate(startDate, endDate)
	var totalProfitLoss float32 = 0.0
	var totalRevenue float32 = 0.0
	var salesCount int = 0
	var productSoldCount int = 0

	for _, srd := range salesReportData {
		totalProfitLoss += srd.ProfitOrLoss
		totalRevenue += srd.TotalSellPricePerProduct
		salesCount += 1
		productSoldCount += srd.Quantity
	}

	var salesReportSummary = reporttypes.SalesReportSummary{
		TotalProfitLoss:  totalProfitLoss,
		TotalRevenue:     totalRevenue,
		SalesCount:       salesCount,
		ProductSoldCount: productSoldCount,
	}

	c.JSON(200, salesReportSummary)
}
