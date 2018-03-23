package main

import (
	"context"
	"flag"

	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"

	"github.com/Synaxis/bfheroesMagma/config"
	"github.com/Synaxis/bfheroesMagma/server"
)

var (
	configFile string
)

func main() {
	initConfig()
	initLogger()

	ctx := context.Background()

	srv, err := server.New()
	if err != nil {
		logrus.Fatal(err)
	}

	srv.ListenAndServe(
		config.Config.HTTPBind,
		config.Config.HTTPSBind,
		config.Config.CertificatePath,
		config.Config.PrivateKeyPath,
	)

	logrus.Info("Listening API requests")
	<-ctx.Done()
	logrus.Info("Exiting...")
}

func initConfig() {
	// Custom path to configuration file
	flag.StringVar(&configFile, "config", ".env", "Path to configuration file")
	flag.Parse()

	// Override env variables
	gotenv.Load(configFile)

	// Initialize config.* public variables
	config.LoadToMemory()
}

func initLogger() {
	logrus.SetLevel(config.LogLevel())

	// logrus.SetFormatter(&logrus.JSONFormatter{
	// 	DisableTimestamp: true,
	// })

	// logrus.SetFormatter(new(prefixed.TextFormatter))
	// logrus.SetFormatter(&prefixed.TextFormatter{
	// 	DisableTimestamp: true,
	// 	DisableColors:    true,
	// })
}
