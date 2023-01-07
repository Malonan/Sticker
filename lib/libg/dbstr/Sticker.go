package dbstr

type Config struct {
	Gid      int64  `gorm:"primaryKey"`
	Data     string `gorm:"type:text(65535);"`
	Admin    string `gorm:"type:text(8000);"`
	Modetype bool
}
