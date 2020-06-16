package database

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type mysqlConfig struct {
	DataSourceName string        `yaml:"dataSourceName"`
	MaxOpenConn    int           `yaml:"maxOpenConn"`
	MaxIdleConn    int           `yaml:"maxIdleConn"`
	ConnMaxLife    time.Duration `yaml:"connMaxLife"`
}

var MysqlConfig mysqlConfig

func initMySQLConfig(file string) error {
	if file == "" {
		return errors.New("mysql config file lost")
	}
	if bytes, err := ioutil.ReadFile(file); err == nil {
		err = yaml.Unmarshal(bytes, &MysqlConfig)
		if err != nil {
			return err
		}
		log.Infof("MySQL config:%v", MysqlConfig)
		return nil
	} else {
		return err
	}
}

//DB mysql数据库连接池，程序结束需close
var DB *sql.DB

func InitDB(file string) error {
	err := initMySQLConfig(file)
	if err != nil {
		log.Errorf("init mysql config error:%v", err)
		return err
	}
	DB, err = sql.Open("mysql", MysqlConfig.DataSourceName)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(MysqlConfig.MaxOpenConn)
	DB.SetMaxIdleConns(MysqlConfig.MaxIdleConn)
	DB.SetConnMaxLifetime(MysqlConfig.ConnMaxLife)
	err = DB.Ping()
	return err
}
