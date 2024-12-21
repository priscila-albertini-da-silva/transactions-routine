package configuration

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type configuration struct {
	Host     string
	Port     int
	LogLevel string
	Database string
}

var Configuration configuration

func setDefaults() {
	viper.SetDefault("Host", "localhost")
	viper.SetDefault("Port", 8080)
	viper.SetDefault("LogLevel", "debug")

	viper.SetDefault("Database", "user='postgres' dbname='transactions_routine' host='127.0.0.1' password='postgres' port='5432' sslmode='disable'")
}

func InitConfig() {
	setDefaults()

	viper.BindEnv("Host", "HOST")
	viper.BindEnv("Port", "PORT")
	viper.BindEnv("LogLevel", "LOGLEVEL")
	viper.BindEnv("Database", "DATABASE_HOST")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(&Configuration); err != nil {
		log.Error("Error Unmarshal:", err)
	}

	logLevel, _ := log.ParseLevel(Configuration.LogLevel)
	log.SetLevel(logLevel)

	log.Info("Initialized configurations")
}
