package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Settings struct {
	ApiKey     string
	ApiSecret  string
	Exchange   string
	Queue      string
	BindingKey string
	lifetime   string
}

func LoadSettings(filenames ...string) *Settings {
	log.Println("LoadSettings")
	err := godotenv.Load(filenames)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Settings{
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
		os.Getenv("EXCHANGE"),
		os.Getenv("QUEUE"),
		os.Getenv("BINDING_KEY"),
		os.Getenv("LIFETIME")}
}
