package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/celsorodrigues/05_func/pkg/routes"
	"github.com/celsorodrigues/05_func/pkg/server"
)

const (
	defaultPort = "3000"
)

func main() {

	data := getData()

	port, ok := os.LookupEnv("SRV_PORT")
	if !ok {
		port = defaultPort
	}
	routes := routes.CustomRoutes(&data)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	server := server.New(
		"localhost",
		port,
		server.WithIdleTimeout(60*time.Second),
		server.WithRoutes(routes),
		server.WithLogger(logger),
	)
	server.Start()
}

func getData() routes.Data {
	return routes.Data{
		Row: []routes.Row{
			{
				Project: "basket",
				CVEs:    []string{"CVE3000"},
			},
			{
				Project: "giftcard",
				CVEs:    []string{"CVE2000"},
			},
		},
	}
}
