package Middleware

import (
	"path"

	"github.com/goccy/go-json"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"

	tele "github.com/3JoB/telebot"
)

var (
	logger *logrus.Logger
)

const (
	LOG_NAME   = "sticker-gramio"
	LOG_SUFFIX = ".log"
	LOG_SIZE   = 60
	LOG_BACKUP = 10
	LOG_DATE   = 7
)

func setOutPut(log *logrus.Logger, log_file_path string) {
	logconf := &lumberjack.Logger{
		Filename:   log_file_path,
		MaxSize:    LOG_SIZE,
		MaxBackups: LOG_BACKUP,
		MaxAge:     LOG_DATE,
		Compress:   true,
	}
	log.SetOutput(logconf)
}

func initLogger() {
	log_file_path := path.Join("./log/", LOG_NAME+LOG_SUFFIX)
	logger = logrus.New()
	setOutPut(logger, log_file_path)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetNoLock()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func Logger() tele.MiddlewareFunc {
	initLogger()
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			data, _ := json.MarshalIndent(c.Update(), "", "  ")
			logger.Println(string(data))
			return next(c)
		}
	}
}
