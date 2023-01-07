package libg

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"sticker/lib/libF"
	log "sticker/lib/liblog"
)

var (
	db *gorm.DB
	kc = libF.F
)

func init() {
	var dialector gorm.Dialector
	dsn := kc().String("database.user") + ":" + kc().String("database.pass") + "@tcp(" + kc().String("database.addr") + ")/" + kc().String("database.db") + "?charset=utf8mb4&parseTime=True&loc=Local"
	dialector = mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	})
	conn, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Use().Println(err.Error())
		print(err)
		return
	}
	sqlDB, err := conn.DB()
	if err != nil {
		log.Use().Println("connect db server failed.")
		print(err)
		return
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(600 * time.Second)

	db = conn
}

func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		log.Use().Println("connect db server failed.")
		print("connect db server failed.")
		return nil
	}
	db.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(d *gorm.DB) {
		d.Statement.RaiseErrorOnNotFound = false
	})
	if err = sqlDB.Ping(); err != nil {
		sqlDB.Close()
	}

	return db
}
