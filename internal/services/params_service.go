package services

import (
	"github.com/gin-gonic/gin"
)

func DecideDate(c *gin.Context) (string, string) {
	queryParams := c.Request.URL.Query()

	var decidedStartDate string
	var decidedEndDate string

	if len(queryParams["start_date"]) == 0 {
		// fmt.Println("empty")
		decidedStartDate = "1990-01-01"
	} else {
		// fmt.Println("no?")
		// fmt.Println(queryParams["start_date"])
		decidedStartDate = queryParams["start_date"][0]
	}

	if len(queryParams["end_date"]) == 0 {
		decidedEndDate = "4000-12-12"
	} else {
		// fmt.Println("no?")
		// fmt.Println(queryParams["end_date"])
		decidedEndDate = queryParams["end_date"][0]
	}

	return decidedStartDate, decidedEndDate
}
