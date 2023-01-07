package bin

import (
	tele "github.com/3JoB/telebot"
)

func Handle(t *tele.Bot) {
	t.Handle("/start", Start)
	t.Handle("/make", StickerBan)
	t.Handle("/refresh", Refresh)
	t.Handle(tele.OnSticker, ServiceSticker)
	t.Handle(tele.OnAddedToGroup, JoinToGroup)
}

func JoinToGroup(c tele.Context) error {
	if c.Chat().Type != "supergroup" {
		return nil
	}
	msg := `Sticker admins have joined the group, are you ready to be filled with my wrath? (just kidding)
You can deploy an identical instance at https://github.com/Malonan/Sticker, or you can learn how to use this bot.`
	fn.S(c, msg)
	Add(c.Chat().ID)
	return GetAdminList(c)
}