package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var lock sync.Mutex

var level string

// Init 初始化
func Init(mode string) {
	level = mode
}

// Info Info日志
func Info(v ...interface{}) {
	writeLog("INFO  ", "", v...)
}

// Infof Info日志
func Infof(format string, args ...interface{}) {
	writeLog("INFO  ", format, args...)
}

// Debug debug日志
func Debug(v ...interface{}) {
	writeLog("DEBUG ", "", v...)
}

// Debugf debug日志
func Debugf(format string, v ...interface{}) {
	writeLog("DEBUG ", format, v...)
}

// Error error日志
func Error(v ...interface{}) {
	writeLog("ERROR ", "", v...)
}

// Errorf error日志
func Errorf(format string, v ...interface{}) {
	writeLog("ERROR ", format, v...)
}

func Fatal(v ...interface{}) {
	writeLog("Fatal ", "", v...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	writeLog("Fatal ", format, v...)
	os.Exit(1)
}

func getFileName() string {
	pwd, _ := os.Getwd()
	fileName := "./" + time.Now().Format("2006-01-02") + ".log"
	return filepath.Join(pwd, "log", fileName)
}

func writeLog(level string, format string, v ...interface{}) {
	var logger *log.Logger
	if level == "release" {
		lock.Lock()
		defer lock.Unlock()
		myfile, err := os.OpenFile(getFileName(), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Printf("Can't open log file:%v", err)
			return
		}
		defer myfile.Close()
		logger = log.New(io.MultiWriter(myfile, os.Stdout), level, log.Ldate|log.Ltime|log.Llongfile)
	} else {
		logger = log.New(os.Stdout, level, log.Ldate|log.Ltime|log.Lshortfile)
	}
	if format == "" {
		logger.Output(3, fmt.Sprintln(v...))
	} else {
		logger.Output(3, fmt.Sprintf(format, v...))
	}
}
