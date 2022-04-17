package models

type (
	EnvConfig struct {
		Port     int            `mapstructure:"port"`
		Database DatabaseConfig `mapstructure:"database"`
	}
	DatabaseConfig struct {
		Timeout int      `mapstructure:"timeout"`
		DBname  string   `mapstructure:"mongo_db_name"`
		URI     []string `mapstructure:"mongo_uri"`
	}
)
