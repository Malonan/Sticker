package log

import (
	"path"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

const (
	LOG_NAME   = "sticker"
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

func InitLogger() {
	log_file_path := path.Join("./log/", LOG_NAME+LOG_SUFFIX)
	logger = logrus.New()
	setOutPut(logger, log_file_path)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetNoLock()
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}


/*func GetLogger() *logrus.Logger {
	return logger
}*/

func Use() *logrus.Logger {
	return logger
}
