package main

import (
	"flag"

	"github.com/musicmash/subscriptions/internal/api"
	"github.com/musicmash/subscriptions/internal/config"
	"github.com/musicmash/subscriptions/internal/db"
	"github.com/musicmash/subscriptions/internal/log"
)

func init() {
	configPath := flag.String("config", "/etc/musicmash/subscriptions/subscriptions.yaml", "Path to subscriptions.yaml config")
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
}

func main() {
	log.Info("Starting subscriptions service...")
	log.Panic(api.ListenAndServe(config.Config.HTTP.IP, config.Config.HTTP.Port))
}
