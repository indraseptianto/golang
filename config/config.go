package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	DBConn string
}

func Load() Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	viper.SetDefault("PORT", "8080")

	return Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}
}
