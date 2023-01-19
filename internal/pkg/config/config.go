package config

type Config struct {
	StorageDir       string `mapstructure:"STORAGE_DIR"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
}
