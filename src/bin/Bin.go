package bin

import (
	"context"

	"github.com/spf13/cast"

	"sticker/lib/libF"
	c "sticker/lib/libc"
	"sticker/lib/libg"
	"sticker/lib/libg/dbstr"
)

var (
	rd = c.Do()       // CacheDB
	db = libg.GetDB() // DataBase
	F  = libF.F()     // Config
	fn Func           // Util

	ctx = context.Background()

	SuperAdmin map[int64]int // SuperAdmin
	WhiteList  map[int64]int // WhiteList
)

func init() {
	SuperAdmin = make(map[int64]int)
	WhiteList = make(map[int64]int)
	Init()
}

func Init() {
	db.AutoMigrate(&dbstr.Config{})

	// Initialize the list of super administrators
	for _, r := range F.Int64s("admin") {
		SuperAdmin[r] = 1
	}

	// Initialize whitelist groups (if enabled)
	if F.Bool("whitelist_mode") {
		for _, r := range F.Int64s("whitelist") {
			WhiteList[r] = 1
		}
	}

	var (
		config []dbstr.Config
	)

	db.Find(&config)

	/*Store database content in Redis

	Please don't use DragonflyDB!!! It causes some problems!!!
	Please wait for DragonflyDB's API version to follow up!!!
	*/
	for _, i := range config {
		rd.Set(ctx, "sticker_Config_Mode_"+cast.ToString(i.Gid), i.Modetype, 0)
		rd.Set(ctx, "sticker_Rule_"+cast.ToString(i.Gid), i.Data, 0)
		rd.Set(ctx, "sticker_Admin_"+cast.ToString(i.Gid), i.Admin, 0)
	}
}
