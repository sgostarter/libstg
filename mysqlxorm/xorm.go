package mysqlxorm

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type DBConfig struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	Debug           bool
	SilentLog       bool
}

func InitXorm(dsn string) (db *xorm.Engine, err error) {
	return InitXormWithConfig(DBConfig{
		DSN: dsn,
	})
}

func InitXormWithConfig(cfg DBConfig) (db *xorm.Engine, err error) {
	engine, err := xorm.NewEngine("mysql", cfg.DSN)
	if err != nil {
		return
	}

	db = engine

	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxIdleConns)
	}

	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	if cfg.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	if cfg.Debug {
		db.ShowSQL(true)
	}

	if cfg.SilentLog {
		db.SetLogLevel(log.LOG_OFF)
	}

	return
}
