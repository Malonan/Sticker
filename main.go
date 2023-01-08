package main

import (
	"sticker/cmd"
	log "sticker/lib/liblog"
)

func init() {
	log.InitLogger()
}

func main() {
	cmd.Start()
}
