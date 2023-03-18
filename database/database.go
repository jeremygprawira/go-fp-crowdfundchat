package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
	DBCharset string
}

/*var Db *gorm.DB
func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}*/

func ConnectDB() (*gorm.DB, error) {
	/*	DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASSWORD"),
		DBHost:    os.Getenv("DB_HOST"),
		DBName:    os.Getenv("DB_NAME"),
		DBCharset: os.Getenv("DB_CHARSET"),
	}*/

	/*DBConfig.DBUser = os.Getenv("DB_USERNAME")
	DBConfig.DBPass = os.Getenv("DB_PASSWORD")
	DBConfig.DBHost = os.Getenv("DB_HOST")
	DBConfig.DBName = os.Getenv("DB_NAME")
	DBConfig.DBCharset = os.Getenv("DB_CHARSET")*/

	/*var (
		DB_USERNAME = os.Getenv("DB_USERNAME")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_HOST = os.Getenv("DB_HOST")
		DB_NAME = os.Getenv("DB_NAME")
		DB_CHARSET = os.Getenv("DB_CHARSET")
	)*/

	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_NAME, DB_CHARSET)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBHost, dbConfig.DBName, dbConfig.DBCharset)
  	dsn := "root:abcd1234@tcp(127.0.0.1:3306)/fp_crowdfundchat_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection is good!")

	return db, nil
}