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

func (u *util) InitMySQLPool(_address, _username, _password, _dbname string, _logSql bool) MySQLPool {
    return &mySQLPool{
        linkString: fmt.Sprintf(
            "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=2s",
            _username, _password, _address, _dbname),
        loggerName: _dbname,
        logSql:     _logSql,
    }
}

func (u *util) InitMySQLPoolByConfig(_mySQLConfig MySQLConfigFormat, logSql bool) MySQLPool {
    return &mySQLPool{
        linkString: fmt.Sprintf(
            "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=2s",
            _mySQLConfig.Username, _mySQLConfig.Password, _mySQLConfig.Address, _mySQLConfig.DBName),
        loggerName: _mySQLConfig.DBName,
        logSql:     logSql,
    }
}

func (wmp *mySQLPool) GetEngine() *xorm.Engine {
    engine, err := xorm.NewEngine("mysql", wmp.linkString)
    if err != nil {
        log.Fatal("Write Database connect error: ", err)
    }
    selfLogger := NewSqlFileLogger(fmt.Sprintf("[%s]", wmp.loggerName), xorm.DEFAULT_LOG_FLAG)
    engine.SetLogger(selfLogger)
    engine.ShowExecTime(wmp.logSql)
    engine.ShowSQL(wmp.logSql)
    return engine
}
