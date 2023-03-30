package core

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type Writer struct {
	logger.Writer
	LogZap bool
}

// NewWriter writer 构造函数
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer, logZap bool) *Writer {
	return &Writer{Writer: w, LogZap: logZap}
}

// Printf 格式化打印日志
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (w *Writer) Printf(message string, data ...interface{}) {
	if w.LogZap {
		zap.L().Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
