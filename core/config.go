package core

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Server struct {
	Port    string
	Debug   bool
	Prefork bool
	TLS     struct {
		Enabled  bool   `mapstructure:"enabled"`
		CertFile string `mapstructure:"cert_file"`
		KeyFile  string `mapstructure:"key_file"`
	}
	Cors struct {
		Enabled bool
	}
	JWT struct {
		Secret string
	}
}

type Database struct {
	Type        string
	Host        string
	Port        int
	User        string
	Password    string
	Name        string
	DBFile      string
	TablePrefix string `mapstructure:"table_prefix"`
	SSLMode     string
	Migrate     bool
}

type Logger struct {
	LogLevel string `mapstructure:"log_level"`
	LogPath  string `mapstructure:"log_path"`
	LogName  string `mapstructure:"log_name"`
	LogExt   string `mapstructure:"log_ext"`
}

type DataHoraderConfig struct {
	Server   Server
	Database Database
	Logger   Logger
}

func defaultDataHoraderConfig() DataHoraderConfig {
	return DataHoraderConfig{
		Server: Server{
			Port:    "11451",
			Debug:   true,
			Prefork: false,
			TLS: struct {
				Enabled  bool   `mapstructure:"enabled"`
				CertFile string `mapstructure:"cert_file"`
				KeyFile  string `mapstructure:"key_file"`
			}{
				Enabled:  false,
				CertFile: "",
				KeyFile:  "",
			},
			Cors: struct {
				Enabled bool
			}{
				Enabled: true,
			},
			JWT: struct {
				Secret string
			}{
				Secret: "data_horader",
			},
		},
		Database: Database{
			Type:        "sqlite3",
			Port:        0,
			DBFile:      "data_horader.db",
			TablePrefix: "sg_",
			Migrate:     true,
		},
		Logger: Logger{
			LogLevel: "debug",
			LogPath:  "logs",
			LogName:  "data_horader",
			LogExt:   "log",
		},
	}
}

func NewDataHoraderConfig() DataHoraderConfig {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.config/.data_horader")
	v.AddConfigPath("/etc/data_horader")
	v.SetConfigName("data_horader")
	v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err = v.SafeWriteConfigAs("data_horader.toml"); err != nil {
				log.Fatalf("Failed to write config: %s", err)
			}
		} else {
			log.Fatalf("Failed to read config: %s", err)
		}
	}

	config := defaultDataHoraderConfig()

	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %s", err)
	}

	return config
}
