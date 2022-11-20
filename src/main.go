package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Runner struct {
	Nonce    int    `json:"nonce"`
	Number   int    `json:"number"`
	Name     string `json:"name"`
	MarketId string `json:"market_id"`
	Close    int    `json:"close"`
	End      int    `json:"end"`
	// Odds Decimal	`json:"odds"`
	PropositionId string `json:"proposition_id"`
}

// func getRunners(context *gin.Context) {
// 	var runners []Runner
// 	if err := context.BindJSON(&runners); err != nil {
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"data": runners})
// }

func getRunners(context *gin.Context) {
	var runners = []Runner{
		{	Nonce: 1,
			Number:        1,
			Name:          "Metaverse",
			MarketId:      "1",
			Close:         1,
			End:           1,
			PropositionId: "1",
		},
	}
	context.IndentedJSON(http.StatusOK, runners)

	// if err := ; err != nil {
	// 	return
	// }

	// context.JSON(http.StatusOK, gin.H{"data": runners})
}

func main() {
	router := gin.Default()

	// router.GET("/runners", func(c *gin.Context) {
	// 	var runners []Runner
	// 	if err := c.BindJSON(&runners); err != nil {
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"data": runners})
	// })

	router.GET("/runners", getRunners)

	router.Run(":8080")
}
