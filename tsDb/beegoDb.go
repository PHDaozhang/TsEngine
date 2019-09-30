package tsDb

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func ConnectDb(driver ...string) error {

	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=" + url.QueryEscape("Local")

	driver_name := "mysql"
	if len(driver) > 0 {
		driver_name = driver[0]

		//dsn = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbuser, dbpassword, dbname, dbhost, dbport)
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbuser, dbpassword, dbhost, dbport, dbname)
	}
	fmt.Println("数据库地址:", dsn)
	err := orm.RegisterDataBase("default", driver_name, dsn, 30, 100)
	if err != nil {
		logs.Debug("数据库服务器链接失败；", err)
	} else {
		logs.Debug("数据库服务器链接成功")
	}

	return err
}

// 自定义传入数据库配置
func ConnectDbWithConfig(appConfig config.Configer, driver ...string) error {

	dbhost := appConfig.String("dbhost")
	dbport := appConfig.String("dbport")
	dbuser := appConfig.String("dbuser")
	dbpassword := appConfig.String("dbpassword")
	dbname := appConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=" + url.QueryEscape("Local")

	driver_name := "mysql"
	if len(driver) > 0 {
		driver_name = driver[0]
		dsn = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbuser, dbpassword, dbname, dbhost, dbport)
	}
	fmt.Println("数据库地址:", dsn)
	err := orm.RegisterDataBase("default", driver_name, dsn)
	if err != nil {
		logs.Debug("数据库服务器链接失败；", err)
	} else {
		logs.Debug("数据库服务器链接成功")
	}

	return err
}

// 采用通用的配置方式
func ConnectDbFormatConfig(driver_name,dbhost,dbport,dbuser,dbpassword,dbname string) error {
	if dbport == "" {
		dbport = "3306"
	}

	dsn := ""
	if driver_name == "mysql" {
		dsn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=" + url.QueryEscape("Local")
	} else {
		dsn = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbuser, dbpassword, dbname, dbhost, dbport)
	}

	fmt.Println("数据库地址:", dsn)
	err := orm.RegisterDataBase("default", driver_name, dsn)
	if err != nil {
		logs.Debug("数据库服务器链接失败；", err)
	} else {
		logs.Debug("数据库服务器链接成功")
	}

	return err
}
