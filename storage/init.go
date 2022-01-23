package storage

import (
	"crud-challenge/config"
	"errors"
	"time"

	"github.com/golobby/container/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	gormConn *gorm.DB
)

func WithStorage() error {
	// this workaround because the docker start MySql with the server in concurrent and MySql need moment for bootstrap
	// In the K8s env, the pod should be terminated and restart in this case.
	db, err := tryConnectStorage()
	if err != nil {
		return errors.New("failed to connect database ERR:" + err.Error())
	}

	gormConn = db

	err = container.Singleton(func() IWagerDAO {
		return WagerDao
	})
	if err != nil {
		return err
	}

	err = container.Singleton(func() IPurchaseDAO {
		return PurchaseDao
	})
	if err != nil {
		return err
	}

	return nil
}

func tryConnectStorage() (db *gorm.DB, err error) {
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(config.Config.MySQL.DSN), &gorm.Config{})
		if err != nil {
			time.Sleep(5*time.Second)
		}
	}

	if err != nil {
		return nil, errors.New("failed to connect database ERR:" + err.Error())
	} else {
		return db, nil
	}
}
