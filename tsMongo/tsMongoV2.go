package tsMongo

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo"
	"time"
)

var MongoClientV2 *mgo.Session

// 初始化MongoDB
func InitMongoV2() {
	db_host := beego.AppConfig.String("Mongodb::db_host")
	auth_db := beego.AppConfig.String("Mongodb::auth_db")
	auth_user := beego.AppConfig.String("Mongodb::auth_user")
	auth_pass := beego.AppConfig.String("Mongodb::auth_pass")
	pool_limit, _ := beego.AppConfig.Int("Mongodb::pool_limit")

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{db_host},
		Timeout:   60 * time.Second,
		Source:    auth_db,
		Username:  auth_user,
		Password:  auth_pass,
		PoolLimit: pool_limit,
	}

	s := &mgo.Session{}
	err := errors.New("")
	if auth_user != "" && auth_pass != "" {
		s, err = mgo.DialWithInfo(dialInfo)
	} else {
		s, err = mgo.Dial(db_host)
	}
	if err != nil {
		logs.Error(fmt.Sprintf("连接MongoDB失败: host:%s, %s\n", db_host, err))
	} else {
		logs.Trace("连接MongoDB成功")
	}

	MongoClientV2 = s
}

func GetMongoDb(db string) *mgo.Database {
	MongoClientV2.SetMode(mgo.Monotonic, true)
	return MongoClientV2.DB(db)
}
