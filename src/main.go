package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var client *http.Client

type Result struct {
	RaceNumber int     `json:"raceNumber"`
	Meeting    Meeting `json:"meeting"`
}

type Meeting struct {
	MeetingName string `json:"meetingName"`
}

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

	url := "https://api.beta.tab.com.au/v1/tab-info-service/racing/dates/2022-11-23/meetings/R/DBN/races/5?jurisdiction=QLD&returnPromo=false"

	// var runners []Runner
	var result Result

	// err := GetJson(url, runners)
	err := GetJson(url, &result)
	if err != nil {
		fmt.Println(err)
	}

	// var runners = []Runner{
	// 	{	Nonce: 1,
	// 		Number:        1,
	// 		Name:          "Metaverse",
	// 		MarketId:      "1",
	// 		Close:         1,
	// 		End:           1,
	// 		PropositionId: "1",
	// 	},
	// }

	context.IndentedJSON(http.StatusOK, result)
}

func GetJson(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	fmt.Println(r.Body)

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
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
