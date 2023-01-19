package config

type Config struct {
	StorageDir    string `mapstructure:"STORAGE_DIR"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}
