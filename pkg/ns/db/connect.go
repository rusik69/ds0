package db

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/sirupsen/logrus"
)

// Connect connects to the database.
func Connect() error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + env.NSEnvInstance.ETCDHost + ":" + env.NSEnvInstance.ETCDPort},
		DialTimeout: 10 * time.Second,
		Username:    env.NSEnvInstance.ETCDUser,
		Password:    env.NSEnvInstance.ETCDPass,
	})
	if err != nil {
		logrus.Println("Error:", err)
		return err
	}
	DB = cli
	return nil
}
