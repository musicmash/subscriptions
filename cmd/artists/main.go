package main

import (
	"flag"

	"github.com/musicmash/artists/internal/config"
	"github.com/musicmash/artists/internal/db"
	"github.com/musicmash/artists/internal/log"
)

func main() {
	configPath := flag.String("config", "/etc/musicmash-artists/artists.yaml", "Path to artists.yaml config")
	flag.Parse()

	if err := config.InitConfig(*configPath); err != nil {
		panic(err)
	}
	if config.Config.Log.Level == "" {
		config.Config.Log.Level = "info"
	}

	log.SetLogFormatter(&log.DefaultFormatter)
	log.ConfigureStdLogger(config.Config.Log.Level)

	db.DbMgr = db.NewMainDatabaseMgr()

	log.Info("Running artists service..")
}
