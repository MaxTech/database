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

type MySQLPool struct {
    linkString string
    loggerName string
    logSql     bool
}

func (u *util) Init(address, username, password, dbname string, logSql bool) *MySQLPool {
    return &MySQLPool{
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

func (u *util) InitByConfig(mySQLConfig MySQLConfigFormat, logSql bool) *MySQLPool {
    return &MySQLPool{
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

func (mp *MySQLPool) GetEngine() *xorm.Engine {
    engine, err := xorm.NewEngine("mysql", mp.linkString)
    if err != nil {
        log.Fatal("Database connect error: ", err)
    }
    engine.ShowExecTime(mp.logSql)
    engine.ShowSQL(mp.logSql)
    engine.SetLogger(
        NewSqlLogger(fmt.Sprintf("[%s]", mp.loggerName), xorm.DEFAULT_LOG_FLAG))
    return engine
}
