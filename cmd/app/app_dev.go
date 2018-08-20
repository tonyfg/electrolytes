// +build !prod

package main

import (
	"net"
	"time"

	"github.com/zserge/webview"
)

func main() {
	// Wait for brunch dev server to boot up (quit trying after ~60s)
	for i := 0; i < 60; i++ {
		conn, err := net.DialTimeout("tcp", "localhost:3333", 1*time.Second)
		if conn != nil {
			conn.Close()
		} else if err != nil {
			time.Sleep(1 * time.Second)
		}
	}

	// Open brunch dev server page inside webview
	window := webview.New(webview.Settings{
		URL:       "http://localhost:3333",
		Width:     800,
		Height:    600,
		Resizable: true,
		Debug:     true,
	})
	window.Run()
}
