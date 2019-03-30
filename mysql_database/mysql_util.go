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

type mySQLPool struct {
    linkString string
    loggerName string
    logSql bool
}

func (u *util) Init(address, username, password, dbname string, logSql bool) *mySQLPool {
    return &mySQLPool{
        linkString: fmt.Sprintf(
            "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=2s",
            username,
            password,
            address,
            dbname),
            loggerName: dbname,
            logSql: logSql,
    }
}

func (mp *mySQLPool) GetEngine() *xorm.Engine {
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
