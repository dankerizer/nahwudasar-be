package routers

import (
	"nahwudasar-be/configs"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
)

func init() {
	configs.ConnectDB()
	conn = configs.DBConn

	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
	})
	log.SetReportCaller(true)
}
