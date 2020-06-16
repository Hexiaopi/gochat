package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/gochat/config/database"
)

func InitAll() error {
	if err := InitLog(); err != nil {
		return err
	}
	if err := database.InitDB("./etc/database/mysql.yaml"); err != nil {
		return err
	}
	return nil
}

func CloseAll(){
	if err:=database.DB.Close();err!=nil{
		log.Errorf("close database error:%v",err)
	}else{
		log.Info("close database success.")
	}
}