package db

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
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"sticker/lib/config"
	log "sticker/lib/log"
)

var (
	db *gorm.DB
	kc = config.Get()
)

func init() {
	var dialector gorm.Dialector
	var err error
	dsn := kc.String("database.user") + ":" + kc.String("database.pass") + "@tcp(" + kc.String("database.addr") + ")/" + kc.String("database.db") + "?charset=utf8mb4&parseTime=True&loc=Local"
	dialector = mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	})
	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Use().Println(err.Error())
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Use().Println("connect db server failed.")
		sqlDB.Close()
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(600 * time.Second)

	db.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(d *gorm.DB) {
		d.Statement.RaiseErrorOnNotFound = false
	})

	go CheckSQLDB(sqlDB)
}

func CheckSQLDB(sqlDB *sql.DB) {
	reconnect := 0
	for {
		if err := sqlDB.Ping(); err != nil {
			if reconnect > 10 {
				sqlDB.Close()
				panic(err)
			}
			reconnect++
		} else {
			if reconnect != 0 {
				reconnect = 0
			} else {
				time.Sleep(time.Minute)
			}
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
