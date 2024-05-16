package postgres_connection

import (
	"fmt"
	"technical_test/config"

	log "github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgres_conn *gorm.DB

func Init() *gorm.DB {
	var err error
	conf := config.Get()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.Database.Host, conf.Database.User, conf.Database.Pass, conf.Database.DatabaseName, conf.Database.Port)
	postgres_conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Error connecting to postgresql: ", err.Error())
		return nil
	}
	log.Info("Sucessfully connection to PostgreSQL")
	return postgres_conn
}

func Get() *gorm.DB {
	if postgres_conn != nil {
		return postgres_conn
	}
	Init()
	return Get()
}
