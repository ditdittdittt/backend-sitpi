package config

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

var (
	Dsn            string
	TimeoutContext time.Duration
	JwtSecret      []byte
)

func Init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	Dsn = fmt.Sprintf("%s?%s", connection, val.Encode())

	TimeoutContext = time.Duration(viper.GetInt("context.timeout")) * time.Second

	JwtSecret = []byte(viper.GetString(`jwt_secret`))
}
