package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/celsorodrigues/05_func/pkg/routes"
	"github.com/celsorodrigues/05_func/pkg/server"
)

const (
	defaultPort = "3000"
)

var (
	//go:embed conf/*
	conf embed.FS
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
	data := routes.Data{}
	confFile, err := conf.ReadFile("conf/conf.json")
	if err != nil {
		fmt.Println("error unmarshaling", err)
	}
	err = json.Unmarshal(confFile, &data)
	if err != nil {
		fmt.Println("cannot unmarshal data struct", err)
	}
	return data
}
