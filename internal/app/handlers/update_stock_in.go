package handlers

import (
	"time"

	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/paramstypes"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/services"
	"github.com/gin-gonic/gin"
)

func UpdateStockIn(c *gin.Context) {
	db := database.DBConn
	var responseCode int
	responseMessage := ""
	var StockInParams paramstypes.StockInCreateWithSku

	if c.ShouldBind(&StockInParams) == nil {
		// =============================================================================
		// VALIDATIONS
		// =============================================================================
		var stockIn = types.StockIn{}

		id := c.Param("id")
		if err := db.First(&stockIn, id).Error; err != nil {
			c.String(404, "Stock In Not Found")
			return
		}

		if StockInParams.ReceivedQuantity > StockInParams.OrderedQuantity {
			c.String(422, "Received Quantity cannot be more than Ordered Quantity")
			return
		}

		var stockInWithSimilarTransactionNumber types.StockIn
		if err := db.Where("transaction_number = ? AND id != ?", StockInParams.TransactionNumber, id).
			Find(&stockInWithSimilarTransactionNumber).Error; err == nil {
			c.String(422, "Transaction Number already exists")
			return
		}

		var product types.Product
		if err := db.Where("barcode = ?", StockInParams.Barcode).First(&product).Error; err != nil {
			c.String(422, "Product does not exist")
			return
		}

		var convertionRate float32 = 1.0
		if len(StockInParams.Currency) > 0 {
			convertionRate = services.ConvertCurrency(StockInParams.Currency)
		}

		// get start date from user or set it up using todays date
		// scuffed i know.
		var startDate time.Time
		var endDate time.Time

		if len(StockInParams.StartTime) > 0 {
			date, _ := time.Parse(time.RFC3339, StockInParams.StartTime+"T00:00:00.000Z")
			startDate = date
		} else {
			startDate = stockIn.StartTime
			endDate = startDate.Add(time.Hour * 24 * time.Duration(product.ExpiresIn))
		}

		if len(StockInParams.EndTime) > 0 {
			date, _ := time.Parse(time.RFC3339, StockInParams.StartTime+"T00:00:00.000Z")
			endDate = date
		} else {
			endDate = stockIn.EndTime
		}

		// =============================================================================

		if errors := db.Model(&stockIn).Updates(&types.StockIn{
			PricePerProduct:   StockInParams.PricePerProduct * convertionRate,
			TransactionNumber: StockInParams.TransactionNumber,
			OrderedQuantity:   StockInParams.OrderedQuantity,
			ReceivedQuantity:  StockInParams.ReceivedQuantity,
			Product:           product,
			Note:              StockInParams.Note,
			StartTime:         startDate,
			EndTime:           endDate,
		}).GetErrors(); len(errors) > 0 {
			responseCode = 422
			for _, err := range errors {
				responseMessage = responseMessage + ", " + err.Error()
			}
		} else {
			responseCode = 200
			responseMessage = "Updated"
		}
	}

	c.String(responseCode, responseMessage)
}
