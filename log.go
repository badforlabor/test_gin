/**
 * Auth :   liubo
 * Date :   2021/9/13 10:13
 * Comment:
 */

package main

import (
	"fmt"
	"github.com/gin-contrib/logger"
	"test_gin/ginlog"

	//"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
)

func mylog(a ...interface{})  {
	//fmt.Fprintln(gin.DefaultWriter, a...)
	log.Info().Msg(fmt.Sprint(a...))
}

var logWriter io.Writer

func initLog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// 修改时间戳格式
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// zerolog务必使用封装好的这个logWriter
	var logConfig = ginlog.GetDefaultLogConfig()
	logWriter = ginlog.Configure(logConfig)

	// 修改zerolog的
	log.Logger = log.Output(logWriter)
	log.Info().
		Bool("fileLogging", logConfig.FileLoggingEnabled).
		Bool("jsonLogOutput", logConfig.EncodeLogsAsJson).
		Str("logDirectory", logConfig.Directory).
		Str("fileName", logConfig.Filename).
		Int("maxSizeMB", logConfig.MaxSize).
		Int("maxBackups", logConfig.MaxBackups).
		Int("maxAgeInDays", logConfig.MaxAge).
		Msg("logging configured")

	// 修改gin的
	gin.DefaultWriter = logWriter
}

func useLogMiddle(router *gin.Engine) {
	router.Use(logger.SetLogger(logger.WithWriter(logWriter)))
}