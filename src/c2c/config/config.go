package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

// Config is a whole config
type Config struct {
	Database DBConfig
}

// DBConfig is config for database
type DBConfig struct {
	Host     string
	User     string
	Password string
	Database string
}

// URL return db url from config
func (dbConfig *DBConfig) URL() string {
	return fmt.Sprintf("%s:%s@%s/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Database)
}

// Get return whole config
func Get() *Config {
	config := Config{}
	_, filename, _, _ := runtime.Caller(0)
	pwd := path.Join(path.Dir(filename))
	configPath := filepath.Join(pwd, "..", "resource", "config.toml")
	log.Println(configPath)
	toml.DecodeFile(configPath, &config)
	return &config
}
