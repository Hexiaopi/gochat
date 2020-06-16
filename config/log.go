package config

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

var loglevel string

func init() {
	flag.StringVar(&loglevel, "log.level", "info", "log level")
}

func InitLog() error {
	level, err := log.ParseLevel(loglevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}
