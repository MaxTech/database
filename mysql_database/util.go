package mysql_database

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "log"
)

type util struct {
}

var Util *util

func init() {
    Util = new(util)
}

type MySQLPool interface {
    GetEngine() *xorm.Engine
}

type mySQLPool struct {
    linkString string
    loggerName string
    logSql     bool
}

func (u *util) InitMySQLPool(address, username, password, dbname string, logSql bool) MySQLPool {
    return &mySQLPool{
        linkString: fmt.Sprintf(
            "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=2s",
            username,
            password,
            address,
            dbname),
        loggerName: dbname,
        logSql:     logSql,
    }
}

func (u *util) InitMySQLPoolByConfig(mySQLConfig MySQLConfigFormat, logSql bool) MySQLPool {
    return &mySQLPool{
        linkString: fmt.Sprintf(
            "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=2s",
            mySQLConfig.Username,
            mySQLConfig.Password,
            mySQLConfig.Address,
            mySQLConfig.DBName),
        loggerName: mySQLConfig.DBName,
        logSql:     logSql,
    }
}

func (wmp *mySQLPool) GetEngine() *xorm.Engine {
    engine, err := xorm.NewEngine("mysql", wmp.linkString)
    if err != nil {
        log.Fatal("Write Database connect error: ", err)
    }
    engine.ShowExecTime(wmp.logSql)
    engine.ShowSQL(wmp.logSql)
    engine.SetLogger(
        NewSqlLogger(fmt.Sprintf("[%s]", wmp.loggerName), xorm.DEFAULT_LOG_FLAG))
    return engine
}
