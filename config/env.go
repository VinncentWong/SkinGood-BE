package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	godotenv.Load()
}

func GetDsn() string {
	dsn := fmt.Sprintf(
		"user=%v password=%v host=%v port=%v dbname=%v",
		os.Getenv("SUPABASE_USER"),
		os.Getenv("SUPABASE_PASSWORD"),
		os.Getenv("SUPABASE_HOST"),
		os.Getenv("SUPABASE_PORT"),
		os.Getenv("SUPABASE_DB_NAME"),
	)
	fmt.Println(dsn)
	return dsn
}
