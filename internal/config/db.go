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
		DbDBName:   getEnv("DB_DBNAME", "sample_db"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbUserName: getEnv("DB_USERNAME", "sample_usr"),
		DbPassword: getEnv("DB_PASSWORD", "sample_pwd"),
	}
}
