package config

type Config struct {
	StorageDir       string `mapstructure:"STORAGE_DIR"`
	ServerHost       string `mapstructure:"SERVER_HOST"`
	ServerPort       string `mapstructure:"SERVER_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
}
