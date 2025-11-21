package config

type Config struct {
	MySQLDSN string
}

func Load() Config {
	return Config{
		MySQLDSN: "root:@tcp(127.0.0.1:3306)/pando?charset=utf8mb4&parseTime=True&loc=Local",
	}
}
