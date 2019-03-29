package mysql_database

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "log"
)

type mySQLUtil struct {
}

var MySQLUtil *mySQLUtil

func (mu *mySQLUtil) InitMySQLEngine(address string, username string, password string, dbname string) *xorm.Engine {
    engine, err := xorm.NewEngine("mysql", fmt.Sprintf(
        "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=2s",
        username,
        password,
        address,
        dbname))
    if err != nil {
        log.Fatal("Database connect error: ", err)
    }
    //defer engine.Close()
    logger = newLogger(fmt.Sprintf("[%s]", dbname), xorm.DEFAULT_LOG_FLAG)
    return engine
}

func (mu *mySQLUtil) CheckMySQLEngine(engine *xorm.Engine) bool {
    err := engine.Ping()
    if err != nil {
        return false
    }
    return true
}
