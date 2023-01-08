package Middleware

/*
  Copyright 2023 Malonan & 3JoB

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

import (
	"path"

	"github.com/goccy/go-json"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"

	tele "github.com/3JoB/telebot"
	"github.com/3JoB/telebot/pkg"
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
			logger.Println(pkg.String(data))
			return next(c)
		}
	}
}
