package handlers

import (
	"fmt"
	"time"

	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/paramstypes"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/services"
	"github.com/gin-gonic/gin"
)

// First check if received quantity > ordered quantity.
// If yes, return 422.
//
// Second find if product by such SKU exists.
// If not return response 422.
//
// If product already exists, add the quantity.
// And create stock in data.
func CreateStockIn(c *gin.Context) {
	db := database.DBConn
	var responseCode int
	responseMessage := ""
	var StockInParams paramstypes.StockInCreateWithSku

	if c.ShouldBind(&StockInParams) == nil {
		// =============================================================================
		// VALIDATIONS
		// =============================================================================
		if StockInParams.ReceivedQuantity > StockInParams.OrderedQuantity {
			c.String(422, "Received Quantity cannot be more than Ordered Quantity")
			return
		}

		var stockInWithSimilarTransactionNumber types.StockIn
		if err := db.Where("transaction_number = ?", StockInParams.TransactionNumber).
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

		// startDate, _ := time.Parse(time.RFC3339, StockInParams.Time+"T00:00:00.000Z")
		var startDate time.Time
		var endDate time.Time

		if len(StockInParams.StartTime) > 0 {
			date, _ := time.Parse(time.RFC3339, StockInParams.StartTime+"T00:00:00.000Z")
			startDate = date
		} else {
			startDate = time.Now()
		}
		if len(StockInParams.EndTime) > 0 {
			fmt.Println("here!")
			date, _ := time.Parse(time.RFC3339, StockInParams.StartTime+"T00:00:00.000Z")
			endDate = date
		} else {
			endDate = time.Now()
			// fmt.Println(time.Hour * 24 * time.Duration(product.ExpiresIn))
			endDate = endDate.AddDate(0, 0, product.ExpiresIn)
		}

		// =============================================================================

		if errors := db.Create(&types.StockIn{
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
			responseCode = 201
			responseMessage = "Created"
		}
	}

	c.String(responseCode, responseMessage)
}
