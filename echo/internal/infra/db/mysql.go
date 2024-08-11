package db

import (
	"fmt"

	"github.com/totoyk/echo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
}

func DBConfig() *Database {
	c := new(Database)

	conf, _ := config.LoadEnv()
	c.Host = conf.DBHost
	c.Port = conf.DBPort
	c.UserName = conf.DBUser
	c.Password = conf.DBPass
	c.DBName = conf.DBName

	return c
}

func MySQL() (*gorm.DB, error) {
	c := DBConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		c.UserName,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
	)
	dbconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dbconn, nil
}
