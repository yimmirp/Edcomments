package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	//Para usar el paquete explicitamente
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Configuration as exported
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

//GetConfiguration as exported
func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//GetConnection as exported
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db

}
