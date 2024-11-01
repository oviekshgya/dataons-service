package config

import (
	"dataons-service/pkg"
	"dataons-service/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

var Config *Configuration

type Configuration struct {
	Server        ServerConfiguration
	DatabaseMysql DatabaseMysqlConfiguration
}

type ServerConfiguration struct {
	AppName string
	Port    string
	Secret  string
	Mode    string
	Env     string
}

type DatabaseMysqlConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

// Setup SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into structs, %v", err)
	}

	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}

type SetDatabaseConfig struct {
	MysqlDb *gorm.DB
}

func setConfiguration(configPath string) {
	Setup(configPath)
	gin.SetMode(GetConfig().Server.Mode)
}

func SetDatabase() *SetDatabaseConfig {
	return &SetDatabaseConfig{
		MysqlDb: SetupMainMysqlDB(),
	}
}

var (
	Conf = GetConfig()
)

func Run() {
	setConfiguration(pkg.CONFIGPATH)
	conf := GetConfig()
	fmt.Println("Pinged your deployment. [ SERVICE IS RUNNING ON", conf.Server.Port, "]")
	Routers := gin.Default()
	setDb := SetDatabase()
	Routers.Use(func(c *gin.Context) {
		c.Set("mysql", setDb.MysqlDb)
		c.Next()
	})
	route.Routers(Routers)
	gin.SetMode(conf.Server.Mode)
	err := Routers.Run(":" + conf.Server.Port)
	if err != nil {
		log.Printf("err running:%s", err.Error())
		return
	}
}
