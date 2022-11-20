package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
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


func main() {
	router := gin.Default()

	router.GET("/runners", func(c *gin.Context) {
		var runners []Runner
		if err := c.BindJSON(&runners); err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": runners})
	})

	router.Run(":8080")
}