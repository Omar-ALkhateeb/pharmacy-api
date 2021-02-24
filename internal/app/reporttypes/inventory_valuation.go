package reporttypes

type InventoryValuation struct {
	ProductBarcode            string
	ProductName               string
	ProductQuantity           int
	ProductAvgPurchasePrice   float32
	ProductTotalPurchasePrice float32
}
