package mysql_database

import (
    "fmt"
    "github.com/go-xorm/xorm"
    "log"
    "os"
    "path/filepath"
    "time"
    "xorm.io/core"
)

// SimpleLogger is the default implment of core.ILogger
type fileLogger struct {
    xorm.SimpleLogger
}

//var _ core.ILogger = &fileLogger{}

func CheckMySQLEngine(_engine *xorm.Engine) bool {
    err := _engine.Ping()
    if err != nil {
        return false
    }
    return true
}

// newLogger let you customrize your logger prefix and flag
func NewSqlFileLogger(_prefix string, _flag int) *fileLogger {
    return initLogger(_prefix, _flag, core.LOG_INFO)
}

// initLogger let you customrize your logger prefix and flag and logLevel
func initLogger(_prefix string, _flag int, _logLevel core.LogLevel) *fileLogger {
    logger := &fileLogger{
        SimpleLogger: xorm.SimpleLogger{
            DEBUG: log.New(nil, fmt.Sprintf("%s [DEBUG] ", _prefix), _flag),
            ERR:   log.New(nil, fmt.Sprintf("%s [ERROR] ", _prefix), _flag),
            INFO:  log.New(nil, fmt.Sprintf("%s [INFO]  ", _prefix), _flag),
            WARN:  log.New(nil, fmt.Sprintf("%s [WARN]  ", _prefix), _flag),
        },
    }
    logger.SimpleLogger.SetLevel(_logLevel)
    return logger
}

// Error implement core.ILogger
func (fl *fileLogger) Error(_v ...interface{}) {
    fl.refreshOutput(fl.ERR)
    fl.SimpleLogger.Error(_v...)
   return
}

// ErrorF implement core.ILogger
func (fl *fileLogger) Errorf(_format string, _v ...interface{}) {
    fl.refreshOutput(fl.ERR)
    fl.SimpleLogger.Errorf(_format, _v...)
   return
}

// Debug implement core.ILogger
func (fl *fileLogger) Debug(_v ...interface{}) {
    fl.refreshOutput(fl.DEBUG)
    fl.SimpleLogger.Debug(_v...)
   return
}

// DebugF implement core.ILogger
func (fl *fileLogger) Debugf(_format string, _v ...interface{}) {
    fl.refreshOutput(fl.DEBUG)
    fl.SimpleLogger.Debugf(_format, _v...)
   return
}

// Info implement core.ILogger
func (fl *fileLogger) Info(_v ...interface{}) {
    fl.refreshOutput(fl.INFO)
    fl.SimpleLogger.Info(_v...)
   return
}

// InfoF implement core.ILogger
func (fl *fileLogger) Infof(_format string, _v ...interface{}) {
    fl.refreshOutput(fl.INFO)
    fl.SimpleLogger.Infof(_format, _v...)
   return
}

// Warn implement core.ILogger
func (fl *fileLogger) Warn(_v ...interface{}) {
    fl.refreshOutput(fl.WARN)
    fl.SimpleLogger.Warn(_v...)
   return
}

// WarnF implement core.ILogger
func (fl *fileLogger) Warnf(_format string, _v ...interface{}) {
    fl.refreshOutput(fl.WARN)
    fl.SimpleLogger.Warnf(_format, _v...)
   return
}

func (fl *fileLogger) refreshOutput(_logger *log.Logger) {
    fileWriter := fl.newWriter()
    _logger.SetOutput(fileWriter)
}

func (fl *fileLogger) newWriter() *os.File {
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
