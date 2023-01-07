package include

import (
	"fmt"
	"time"

	tele "github.com/3JoB/telebot"
	"github.com/3JoB/telebot/middleware"

	"sticker/lib/libF"
	fn "sticker/lib/libfn"
	log "sticker/lib/liblog"
	m "sticker/src/Middleware"
	"sticker/src/bin"
)

var F = libF.F()

func T() string {
	return "[Runtime/" + time.Now().UTC().Format("2006-01-02 15:04:05") + "]"
}

func Start() {
	fmt.Println(T() + " Ready to start.....")
	fmt.Println(T() + " Configure robot information...")
	pref := tele.Settings{
		Token:   F.String("Key"),
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, ctx tele.Context) { log.Use().Error(err.Error()) },
		Client:  fn.Client(),
	}
	fmt.Println(T() + " Registering...")
	b, err := tele.NewBot(pref)
	if err != nil {
		fmt.Println(T() + " Registration failed, please check the log.")
		log.Use().Println(err.Error())
		return
	}
	b.RemoveWebhook(true)
	b.Use(middleware.Recover())
	// b.Use(middleware.AutoRespond())
	b.Use(m.Logger())
	bin.Handle(b)
	fmt.Println(T() + " The robot is running on @" + b.Me.Username)
	b.Start()
}
