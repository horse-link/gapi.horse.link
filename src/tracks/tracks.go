package tracks

import (
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// type Margin struct {
// 	MeetingName string `json:"meetingName"`
// 	Location    string `json:"location"`
// }

type Profile struct {
	Name      string
	Nmemmonic string
	Enabled   bool
}

func GetTracks(context *gin.Context) {
	// https://bun.uptrace.dev/postgres/#pgdriver
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
