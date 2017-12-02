package config

import (
	"fmt"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/viper"
)

// DSN stores all the database connection and driver information
type DSN struct {
	Driver             string
	Host               string
	Port               string
	Name               string
	Username           string
	Password           string
	ParseTime          string
	MaxOpenConnections int
	MaxIdleConnections int
}

func getHomePath() string {
	home, _ := homedir.Dir()
	return home
}

func getProjectPath() string {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Warn("Warning, cannot get current path")
		return ""
	}
	// Traverse back from current directory until service base dir is reaach and add to config path
	for !strings.HasSuffix(dir, "go-boilerplate") && dir != "/" {
		dir, err = filepath.Abs(dir + "/..")
		if err != nil {
			break
		}
	}
	return dir
}

// Init loads up the application config file
func Init() {
	// Find home directory.
	viper.SetEnvPrefix("go-boilerplate")
	viper.BindEnv("configFile")
	viper.BindEnv("configPath")

	viper.SetDefault("configFile", "config")
	viper.SetDefault("configPath", "/tmp")

	viper.AddConfigPath(getHomePath())
	viper.AddConfigPath(getProjectPath())
	viper.AddConfigPath(viper.GetString("configPath"))
	viper.SetConfigName(viper.GetString("configFile"))

	viper.SetDefault("logging.level", "DEBUG")
	viper.SetDefault("logging.errorLogFile", "error.log")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Errorln("Error using config file:", viper.ConfigFileUsed())
		log.Errorln(err.Error())
	}
}

// ConnectionString returns the correctly formatted connection string for connecting to the database
func (d DSN) ConnectionString() string {
	return fmt.Sprintf("./%s.db", d.Name)
}

// GetDSN : returns the database dsn from viper config.
func GetDSN() DSN {
	dsn := DSN{Driver: viper.GetString("store.db.driver"),
		Username:           viper.GetString("store.db.username"),
		Password:           viper.GetString("store.db.password"),
		Host:               viper.GetString("store.db.host"),
		Port:               viper.GetString("store.db.port"),
		Name:               viper.GetString("store.db.name"),
		ParseTime:          viper.GetString("store.db.parseTime"),
		MaxOpenConnections: viper.GetInt("store.db.maxOpenConnections"),
		MaxIdleConnections: viper.GetInt("store.db.maxIdleConnections")}
	return dsn
}
