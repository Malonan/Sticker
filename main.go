package main

import (
	"sticker/cmd"
	"sticker/lib/log"
)

func init() {
	log.InitLogger()
}

func main() {
	cmd.Start()
}
