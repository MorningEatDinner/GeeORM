package log

import (
	"io"
	"log"
	"os"
	"sync"
)

/*
log.Lshortfile 是一个标志位，用于在日志消息中包含文件名和行号的短格式。
log.LstdFlags 是一个预定义的标志位，包括日期、时间等标准信息。
使用 log.Lshortfile 支持显示文件名和代码行号。
*/

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info]\033[0m", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Print
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// 设置日志的等级
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	// 如果当下设计的等级要高于Error，那就是discard了， 那么就是将 errorLog 日志记录器的输出重定向到 ioutil.Discard，丢弃
	if level > ErrorLevel {
		errorLog.SetOutput(io.Discard) // 如果
	}
	// 如果是ErrorLevel或者Disabled等级的设置， 那么低级别的info就不能够输出
	if level > InfoLevel {
		infoLog.SetOutput(io.Discard)
	}
}
