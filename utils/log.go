package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var Log *Logger

type Logger struct {
	fileDate string
	file     *os.File
	trace    *log.Logger
	info     *log.Logger
	warning  *log.Logger
	error    *log.Logger
}

func InitLog() *os.File {
	splitDate := getSplitDate()
	file := newFile(splitDate)
	logFormat := log.Ldate | log.Ltime | log.Llongfile | log.LstdFlags
	logger := Logger{
		fileDate: splitDate,
		file:     file,
		trace:    log.New(io.MultiWriter(file, os.Stderr), "[TRACE] ", logFormat),
		info:     log.New(io.MultiWriter(file, os.Stderr), "[INFO] ", logFormat),
		warning:  log.New(io.MultiWriter(file, os.Stderr), "[WARNING] ", logFormat),
		error:    log.New(io.MultiWriter(file, os.Stderr), "[ERROR] ", logFormat),
	}
	Log = &logger
	return Log.file
}

func getSplitDate() string {
	splitType := os.Getenv("LOG_SPLIT_TYPE")
	switch splitType {
	case "hour":
		return time.Now().Format("2006-01-02_15")
	case "day":
		return time.Now().Format("2006-01-02")
	case "week":
		return GetFirstDateOfWeek()
	case "month":
		return time.Now().Format("2006-01")
	default:
		return time.Now().Format("2006-01-02")
	}
}

func newFile(name string) *os.File {
	file, err := os.OpenFile(
		os.Getenv("LOG_PATH")+"/"+name+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file：%s", err))
	}
	return file
}

func (l *Logger) checkSplitFile() {
	splitDate := getSplitDate()
	if l.fileDate != splitDate {
		if err := l.file.Close(); err != nil {
			l.warning.Println("Failed to close log file:", err)
		}
		l.fileDate = splitDate
		l.file = newFile(l.fileDate)
		l.trace.SetOutput(io.MultiWriter(l.file, os.Stderr))
		l.info.SetOutput(io.MultiWriter(l.file, os.Stderr))
		l.warning.SetOutput(io.MultiWriter(l.file, os.Stderr))
		l.error.SetOutput(io.MultiWriter(l.file, os.Stderr))
	}
}

func (l *Logger) Trace(log ...interface{}) {
	l.checkSplitFile()
	l.trace.Println(log...)
}

func (l *Logger) Info(log ...interface{}) {
	l.checkSplitFile()
	l.info.Println(log...)
}

func (l *Logger) Warning(log ...interface{}) {
	l.checkSplitFile()
	l.warning.Println(log...)
}

func (l *Logger) Error(log ...interface{}) {
	l.checkSplitFile()
	l.error.Println(log...)
}
