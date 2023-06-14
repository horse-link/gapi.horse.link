package market

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// type Margin struct {
// 	MeetingName string `json:"meetingName"`
// 	Location    string `json:"location"`
// }

func GetMargin(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}