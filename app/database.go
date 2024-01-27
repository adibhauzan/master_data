package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	//config := viper.New()
	//config.SetConfigName("config")
	//config.SetConfigType("json")
	//config.AddConfigPath("../")
	//
	//err := config.ReadInConfig()
	//if err != nil {
	//	panic(err)
	//}
	//
	//dbHost := config.GetString("database.host")
	//dbUser := config.GetString("database.user")
	//dbPassword := config.GetString("database.password")
	//dbName := config.GetString("database.name")
	//dbPort := config.GetInt("database.port")
	//dbSslMode := config.GetString("database.ssl_mode")
	//dbTimeZone := config.GetString("database.time_zone")
	//
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode, dbTimeZone)
	dsn := "host=localhost user=postgres password=adib5665 dbname=master_data port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}
