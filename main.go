package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/gochat/api"
	"github.com/hexiaopi/gochat/config"
)

const PIDFileName = ".pid.pid"

func main() {
	flag.Parse()

	if err := config.InitAll(); err != nil {
		log.Errorf("init config fail :%v", err)
		return
	}
	_ = ioutil.WriteFile(PIDFileName, []byte(fmt.Sprintf("%d", os.Getpid())), os.ModeType)

	go func() { api.StartServer() }()

	killerChan := make(chan os.Signal, 1)
	signal.Notify(killerChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case sig := <-killerChan:
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		log.Infof("get killer signal. %v", sig)
		api.StopServer(ctx)
		//关闭数据库连接
		config.CloseAll()
		_ = os.Remove(PIDFileName)
		log.Info("pid remove success.")
		return
	}
}
