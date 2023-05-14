package config

import "github.com/spf13/viper"

type Config struct {
	// Mysql setup
	DBHost     string  `mapstructure:"SQL_HOST"`
	DBUsername string  `mapstructure:"SQL_USER"`
	DBPassword string  `mapstructure:"SQL_PASSWORD"`
	DBName     string  `mapstructure:"SQL_DB"`
	DBPORT     string  `mapstructure:"SQL_PORT"`
	
	
	// Redis Setup
	RedisUrl   string  `mapstructure:"REDIS_URL"`
}


func LoadConfig(path string)(config Config, err error) {

	viper.addConfigPath(path)
	viper.SetupConfigType("env")
	viper.SetupConfigName("env")
	
	viper.AutomaticEnv()
	
	// handle null
	err = viper.ReadInConfig()
	if(err != nil){
		return 
	}
	
	err = viper.Unmarshal(&config)
	return
	
}