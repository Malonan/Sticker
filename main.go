package main

import (
	"sticker/include"
	log "sticker/lib/liblog"
)

func init() {
	log.InitLogger()
}

func main() {
	include.Start()
}
