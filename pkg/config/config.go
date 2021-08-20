package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// Server - struct for server config
type Server struct {
	Port string `toml:"port"`
}

// Database - struct for Database config
type Database struct {
	ConnectionString string `toml:"connection_string"`
	DatabaseName     string `toml:"database_name"`
}

// Jwt - struct for JWT tokens konfigs
type Jwt struct {
	JwtKey string `toml:"key"`
}

// Logs - struct for Logs configs
type Logs struct {
	LogFilePath string `toml:"path"`
}

// Config  - Main struct for all settings
type Config struct {
	ServerConfig   Server   `toml:"server"`
	DatabaseConfig Database `toml:"database"`
	JwtConfig      Jwt      `toml:"jwt"`
	LogsConfig     Logs     `toml:"logs"`
}

// ResultConfig - Result struct for store settings
var ResultConfig Config

// ApplyDefaultConfig - Set Default settings if config file not exists or can't read file
func ApplyDefaultConfig(cfg *Config) {

	//set server default config
	cfg.ServerConfig.Port = ":8000"
	//set DB default config
	cfg.DatabaseConfig.ConnectionString = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
	cfg.DatabaseConfig.DatabaseName = "test"
	//set JWT default config
	cfg.JwtConfig.JwtKey = "aseredqwefs123ddgaqwetsdg465"
}

//ParseConfig - use for parse toml config file and set default settings if file not found
func ParseConfig(configFile string) {

	var conf Config
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		fmt.Println("File not found\tSet up default settings")
		ApplyDefaultConfig(&ResultConfig)
		return
	}
	fmt.Println("cfg\n:", conf)
	ResultConfig = conf

}
