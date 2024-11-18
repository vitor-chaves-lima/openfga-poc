package main

import "os"

type Config struct {
	OpenFgaUrl string
}

func GetConfig() Config {
	OpenFgaUrl := os.Getenv("OPEN_FGA_URL")
	if OpenFgaUrl == "" {
		panic("OPEN_FGA_URL is not set")
	}

	return Config{
		OpenFgaUrl,
	}
}
