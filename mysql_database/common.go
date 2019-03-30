package mysql_database

import (
    "fmt"
    "github.com/go-xorm/core"
    "github.com/go-xorm/xorm"
    "log"
    "os"
    "path/filepath"
    "time"
)

// SimpleLogger is the default implment of core.ILogger
type dbLogger struct {
    DEBUG   *log.Logger
    ERR     *log.Logger
    INFO    *log.Logger
    WARN    *log.Logger
    level   core.LogLevel
    showSQL bool
}

var _ core.ILogger = &dbLogger{}

func CheckMySQLEngine(engine *xorm.Engine) bool {
    err := engine.Ping()
    if err != nil {
        return false
    }
    return true
}

// newLogger let you customrize your logger prefix and flag
func NewSqlLogger(prefix string, flag int) *dbLogger {
    return initLogger(prefix, flag, core.LOG_INFO)
}

// initLogger let you customrize your logger prefix and flag and logLevel
func initLogger(prefix string, flag int, l core.LogLevel) *dbLogger {
    return &dbLogger{
        DEBUG: log.New(nil, fmt.Sprintf("%s [debug] ", prefix), flag),
        ERR:   log.New(nil, fmt.Sprintf("%s [error] ", prefix), flag),
        INFO:  log.New(nil, fmt.Sprintf("%s [info]  ", prefix), flag),
        WARN:  log.New(nil, fmt.Sprintf("%s [warn]  ", prefix), flag),
        level: l,
    }
}

// Error implement core.ILogger
func (s *dbLogger) Error(v ...interface{}) {
    if s.level <= core.LOG_ERR {
        fileWriter := s.newWriter()
        s.ERR.SetOutput(fileWriter)
        defer fileWriter.Close()
        s.ERR.Output(2, fmt.Sprint(v...))
    }
    return
}

// Errorf implement core.ILogger
func (s *dbLogger) Errorf(format string, v ...interface{}) {
    if s.level <= core.LOG_ERR {
        fileWriter := s.newWriter()
        s.ERR.SetOutput(fileWriter)
        defer fileWriter.Close()
        s.ERR.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Debug implement core.ILogger
func (s *dbLogger) Debug(v ...interface{}) {
    if s.level <= core.LOG_DEBUG {
        fileWriter := s.newWriter()
        s.DEBUG.SetOutput(fileWriter)
        defer fileWriter.Close()
        s.DEBUG.Output(2, fmt.Sprint(v...))
    }
    return
}

// Debugf implement core.ILogger
func (s *dbLogger) Debugf(format string, v ...interface{}) {
    if s.level <= core.LOG_DEBUG {
        fileWriter := s.newWriter()
        s.DEBUG.SetOutput(fileWriter)
        defer fileWriter.Close()
        s.DEBUG.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Info implement core.ILogger
func (s *dbLogger) Info(v ...interface{}) {
    if s.level <= core.LOG_INFO {
        fileWriter := s.newWriter()
        s.INFO.SetOutput(fileWriter)
        defer fileWriter.Close()
        s.INFO.Output(2, fmt.Sprint(v...))
    }
    return
}

// Infof implement core.ILogger
func (s *dbLogger) Infof(format string, v ...interface{}) {
    if s.level <= core.LOG_INFO {
        fileWriter := s.newWriter()
        s.INFO.SetOutput(fileWriter)
        defer fileWriter.Close()
        s.INFO.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Warn implement core.ILogger
func (s *dbLogger) Warn(v ...interface{}) {
    if s.level <= core.LOG_WARNING {
        fileWriter := s.newWriter()
        s.WARN.SetOutput(fileWriter)
        defer fileWriter.Close()
        s.WARN.Output(2, fmt.Sprint(v...))
    }
    return
}

// Warnf implement core.ILogger
func (s *dbLogger) Warnf(format string, v ...interface{}) {
    if s.level <= core.LOG_WARNING {
        fileWriter := s.newWriter()
        s.WARN.SetOutput(fileWriter)
        defer fileWriter.Close()
        s.WARN.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Level implement core.ILogger
func (s *dbLogger) Level() core.LogLevel {
    return s.level
}

// SetLevel implement core.ILogger
func (s *dbLogger) SetLevel(l core.LogLevel) {
    s.level = l
    return
}

// ShowSQL implement core.ILogger
func (s *dbLogger) ShowSQL(show ...bool) {
    if len(show) == 0 {
        s.showSQL = true
        return
    }
    s.showSQL = show[0]
}

// IsShowSQL implement core.ILogger
func (s *dbLogger) IsShowSQL() bool {
    return s.showSQL
}

func (s *dbLogger) newWriter() *os.File {
    path, _ := filepath.Abs("./logs/db")
    _, err := os.Stat(path)
    if err != nil && os.IsNotExist(err) {
        _ = os.MkdirAll(path, os.ModePerm)
    }

    logFile := fmt.Sprintf("./logs/db/db_%s.log", time.Now().Format("20060102"))
    logFile, _ = filepath.Abs(logFile)
    fileWriter, _ := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
    return fileWriter
}
