package configs

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "900")
	viper.SetDefault("datanase.host", "localhost")
	viper.SetDefault("database.port", "5432")

}

func Load() error {
	viper.SetconfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nill {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

}
