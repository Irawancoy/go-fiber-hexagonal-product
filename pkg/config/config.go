package config

type Config struct {
	ServerAddress     string
	MongoURI          string
	MySQLDSN          string
	MongoDatabaseName string
}

func LoadConfig() *Config {
	return &Config{
		ServerAddress:     ":8080",
		MySQLDSN:          "root:mysql123456@tcp(localhost:3306)/goproduct_db",
		MongoURI:          "mongodb://localhost:27017",
		MongoDatabaseName: "goproduct_db",
	}
}
