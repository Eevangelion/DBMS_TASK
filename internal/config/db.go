package config

type Database struct {
	DbDBName   string
	DbHost     string
	DbPort     string
	DbUserName string
	DbPassword string
}

func SetupDB() *Database {
	return &Database{
		DbDBName:   getEnv("DB_DBNAME", "FunnyJokesDB"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbUserName: getEnv("DB_USERNAME", "postgres"),
		DbPassword: getEnv("DB_PASSWORD", "root"),
	}
}
