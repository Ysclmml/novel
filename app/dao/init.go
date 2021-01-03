package dao

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"log"
	"novel/app/global/my_errors"
	zeroLog "novel/app/log"
	"strings"
	"time"
)

var (
	db *gorm.DB
)

// 当前支持的数据库类型
const (
	DriverMysql  = "mysql"
	DriverSqlServer = "sqlserver"
	DriverPostgresql = "postgresql"
)

func SetUp() {
	if dbDriver, err := getSqlDriver(); err != nil {
		log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
	} else {
		db = dbDriver
	}
}

// 获取sql驱动
func getSqlDriver() (*gorm.DB, error) {
	var dbDialect gorm.Dialector

	sqlType := viper.GetString("database.sqlType")

	if val, err := getDbDialect(sqlType, "write"); err != nil {
		zeroLog.Error(my_errors.ErrorsDialectorDbInitFail + err.Error())
	} else {
		dbDialect = val
	}
	// gorm配置可以参考 https://gorm.io/zh_CN/docs/gorm_config.html
	gormDb, err := gorm.Open(dbDialect, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		// todo: 日志待研究
		// Logger:                 redefineLog(sqlType), //拦截、接管 gorm v2 自带日志
	})
	if err != nil {
		// gorm 数据库驱动初始化失败
		return nil, err
	}

	// 如果开启了读写分离，配置读数据库（resource、read、replicas）
	if viper.GetInt("database."+sqlType + ".isOpenReadDb") == 1 {
		if val, err := getDbDialect(sqlType, "read"); err != nil {
			zeroLog.Error(my_errors.ErrorsDialectorDbInitFail + err.Error())
		} else {
			dbDialect = val
		}
		resolverConf := dbresolver.Config{
			Replicas: []gorm.Dialector{dbDialect}, //  读 操作库，查询类
			Policy:   dbresolver.RandomPolicy{},     // sources/replicas 负载均衡策略适用于
		}
		err = gormDb.Use(dbresolver.Register(resolverConf).SetConnMaxIdleTime(time.Second * 30).
			SetConnMaxLifetime(viper.GetDuration("database." + sqlType +".read.setConnMaxLifetime") * time.Second).
			SetMaxIdleConns(viper.GetInt("database." + sqlType + ".read.setMaxIdleConns")).
			SetMaxOpenConns(viper.GetInt("database." + sqlType + ".read.setMaxOpenConns")))
		if err != nil {
			return nil, err
		}
	}

	// 查询没有数据，屏蔽 gorm v2 包中会爆出的错误
	// https://github.com/go-gorm/gorm/issues/3789  此 issue 所反映的问题就是我们本次解决掉的
	_ = gormDb.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(d *gorm.DB) {
		d.Statement.RaiseErrorOnNotFound = false
	})

	// 为主连接设置连接池(48行返回的数据库驱动指针)
	if rawDb, err := gormDb.DB(); err != nil {
		return nil, err
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(viper.GetDuration("database."+ sqlType + ".write.setConnMaxLifetime") * time.Second)
		rawDb.SetMaxIdleConns(viper.GetInt("database." + sqlType + ".write.setMaxIdleConns"))
		rawDb.SetMaxOpenConns(viper.GetInt("database." + sqlType + ".write.setMaxOpenConns"))
		return gormDb, nil
	}
}

// 获取一个数据库方言(Dialector),通俗的说就是根据不同的连接参数，获取具体的一类数据库的连接指针
func getDbDialect(sqlType, readWrite string) (gorm.Dialector, error) {
	var dbDialect gorm.Dialector
	dsn := getDsn(sqlType, readWrite)
	switch strings.ToLower(sqlType) {
	case DriverMysql:
		dbDialect = mysql.Open(dsn)
	case DriverSqlServer:
		dbDialect = sqlserver.Open(dsn)
	case DriverPostgresql:
		dbDialect = postgres.Open(dsn)
	default:
		return nil, errors.New(my_errors.ErrorsDbDriverNotExists + sqlType)
	}
	return dbDialect, nil
}

// 根据配置参数生成数据库驱动 dsn
func getDsn(sqlType, readWrite string) string {
	host := viper.GetString(fmt.Sprintf("database.%s.%s.host", sqlType, readWrite))
	user := viper.GetString(fmt.Sprintf("database.%s.%s.user", sqlType, readWrite))
	password := viper.GetString(fmt.Sprintf("database.%s.%s.password", sqlType, readWrite))
	name := viper.GetString(fmt.Sprintf("database.%s.%s.name", sqlType, readWrite))
	port := viper.GetInt(fmt.Sprintf("database.%s.%s.port", sqlType, readWrite))
	charset := viper.GetString(fmt.Sprintf("database.%s.%s.charset", sqlType, readWrite))

	switch strings.ToLower(sqlType) {
	case DriverMysql:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, name, charset)
	case DriverSqlServer:
		return fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable", host, port, name, user, password)
	case DriverPostgresql:
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, name, user, password)
	}
	return ""
}

func GetDb() *gorm.DB {
	return db
}