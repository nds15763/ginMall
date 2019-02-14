package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logrus *logrus.Entry
}

//logger方法我没具体写，有兴趣的朋友按需实现吧
func NewLogger() *Logger {
	logger := &Logger{
		logrus: logrus.WithFields(logrus.Fields{}),
	}
	return logger
}

func (logger *Logger) Info(str error) {
	fmt.Print("Info：" + str.Error())
}

func (logger *Logger) Warn(str error) {
	fmt.Print("Warn：" + str.Error())
}

func (logger *Logger) Error(str error) {
	fmt.Print("Error：" + str.Error())
}
