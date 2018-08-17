package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/zserge/webview"
)

type config struct {
	Env string `env:"MY_APP_ENV" envDefault:"dev"`
}

func main() {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file - %v", err)
		}
	} else {
		log.Print(".env file not found. Using default configuration...")
	}

	cfg := config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Error setting configuration: %+v", err)
	}
	fmt.Printf("Configuration: %+v\n", cfg)

	if cfg.Env == "dev" {
		// Wait for brunch dev server to boot up (quit trying after ~10s)
		for i := 0; i < 10; i++ {
			conn, err := net.DialTimeout("tcp", "localhost:3333", 1*time.Second)
			if conn != nil {
				conn.Close()
			} else if err != nil {
				time.Sleep(1 * time.Second)
			}
		}

		// Open webpack dev server page inside webview
		window := webview.New(webview.Settings{
			URL:       "http://localhost:3333",
			Width:     800,
			Height:    600,
			Resizable: true,
			Debug:     true})
		window.Run()
	} else {
		// Open the compiled build inside webview
		webview.Open(
			"WIP WIP WIP",
			"wip wip wip",
			800,
			600,
			true)
	}
}
