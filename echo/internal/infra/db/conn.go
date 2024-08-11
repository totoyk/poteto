package db

import (
	"fmt"

	"github.com/totoyk/echo/config"
	"gorm.io/gorm"
)

func NewConn() (dbconn *gorm.DB, err error) {
	conf, _ := config.LoadEnv()

	switch conf.DBType {
	case "mysql":
		dbconn, err = MySQL()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid db type: %s", conf.DBType)
	}

	return
}
