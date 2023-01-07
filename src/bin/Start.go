package bin

import (
	tele "github.com/3JoB/telebot"
	"github.com/spf13/cast"
)

func Start(c tele.Context) error {
	if c.Chat().Type != "private" {
		return nil
	}
	c.Send("Hello World!")
	return nil
}

func Refresh(c tele.Context) error {
	if c.Chat().Type != "supergroup" {
		c.Send("What?")
		return nil
	}
	admin := Int64Map(rd.Get(ctx, "sticker_Admin_"+cast.ToString(c.Chat().ID)).Result())
	if len(admin) == 0 {
		GetAdminList(c)
		fn.SA(c, 10, "The current group management list is empty and is trying to get it.\nIf you can't get it, check that the bot has been granted administrator privileges.")
		return nil
	}
	if admin[c.Sender().ID] != 1 {
		fn.SA(c, 10, "This command is only available to supergroup administrators!!!")
		return nil
	}
	fn.SA(c, 6, "Refreshing admin list...")
	return GetAdminList(c)
}
