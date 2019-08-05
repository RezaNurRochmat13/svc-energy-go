package resources

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

func SetupDatabaseConnection() (*sql.DB, error) {
	viper.SetConfigName("mysql")
	viper.AddConfigPath("environment/development")

	handleReadingConfigFile := viper.ReadInConfig()

	if handleReadingConfigFile != nil {
		return nil, handleReadingConfigFile
	}

	var MySQLHost = viper.GetString("mysql.MySQLHost")
	var MySQLUsername = viper.GetString("mysql.MySQLUsername")
	var MySQLPassword = viper.GetString("mysql.MySQLPassword")
	var MySQLDatabaseName = viper.GetString("mysql.MySQLDatabaseName")

	var databaseConfig = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		MySQLUsername, MySQLPassword, MySQLHost, MySQLDatabaseName)

	databaseConnectionConfiguration, errorDatabaseConfiguration := sql.Open("mysql", databaseConfig)

	if errorDatabaseConfiguration != nil {
		return nil, errorDatabaseConfiguration
	}

	databaseConnectionConfiguration.SetMaxOpenConns(50)
	databaseConnectionConfiguration.SetMaxIdleConns(50)

	return databaseConnectionConfiguration, nil

}
