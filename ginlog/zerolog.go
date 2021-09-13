/**
 * Auth :   liubo
 * Date :   2021/9/13 13:40
 * Comment:
 */

package ginlog

import (
	"fmt"
	"github.com/badforlabor/gocrazy/crazyos"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path"
	"time"
)

// Configuration for logging
type LogConfig struct {
	// Enable console logging
	ConsoleLoggingEnabled bool
	PrettyConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

func GetDefaultLogConfig() *LogConfig {

	var logfilename = crazyos.GetAppName()
	var logDir = ""

	logDir = crazyos.GetExecFolder()
	logDir = path.Join(logDir, "logs")

	var cfg = LogConfig { ConsoleLoggingEnabled:true, EncodeLogsAsJson:false,
		FileLoggingEnabled:true, Directory:logDir, Filename:logfilename,
		MaxBackups:100,MaxAge:7, MaxSize:32}


	return &cfg
}


var svcIsWindowsService func()(bool, error)

func Configure(config *LogConfig) io.Writer {
	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		var ws = false
		if svcIsWindowsService != nil {
			ws, _ = svcIsWindowsService()
		}
		if !ws {
			writers = append(writers, os.Stdout)
		}
	}
	if config.PrettyConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}
	mw := io.MultiWriter(writers...)

	return mw
}

func newRollingFile(config *LogConfig) io.Writer {
	if err := os.MkdirAll(config.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", config.Directory).Msg("can't create log directory")
		return nil
	}

	var postfix = time.Now().Format("20060102-150405")
	//postfix = "%Y%m%d%H%M"

	var rot, e = rotatelogs.New(path.Join(config.Directory, config.Filename + "." + postfix + ".log"),
		rotatelogs.WithRotationSize(int64(config.MaxSize) * 1024 * 1024),
		rotatelogs.ForceNewFile(),
	)
	if e != nil {
		fmt.Println("create log failed：", e)
	}
	return rot

	//return &lumberjack.Logger{
	//	Filename:   path.Join(config.Directory, config.Filename),
	//	MaxBackups: config.MaxBackups,		// files
	//	MaxSize:    config.MaxSize,			// megabytes
	//	MaxAge:     config.MaxAge,			// days
	//}
}
