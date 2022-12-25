package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

var DB *gorm.DB

const (
	DB_HOST     = "127.0.0.1"
	DB_DRIVER   = "postgres"
	DB_USER     = "root"
	DB_PASSWORD = "secret"
	DB_NAME     = "hackernews"
	DB_PORT     = "5432"
)

func LoadConfig() (config Config, err error) {
	config.DBDriver = os.Getenv("DB_DRIVER")
	if config.DBDriver == "" {
		config.DBDriver = DB_DRIVER
	}
	if os.Getenv("DB_HOST") != "" {
		config.DBSource = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	} else {
		config.DBSource = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)
	}
	return config, err
}

func NewDB(conf Config) *gorm.DB {
	DB, err := gorm.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}
	return DB
}
