package Config
// Db Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


var DB *gorm.DB


// represent DB config
type DBConfig struct {
	Driver string
	Host string
	Port int
	User string
	DBName string
	Password string
}


func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Driver: "postgresql",
		Host: "localhost",
		Port: 5431,
		User: "postgres",
		DBName: "go_db",
		Password: "password",
	}
	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s",
		dbConfig.Driver,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}