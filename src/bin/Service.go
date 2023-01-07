package bin

import (
	tele "github.com/3JoB/telebot"
	"github.com/spf13/cast"
)

func ServiceSticker(c tele.Context) error {
	if c.Chat().Type != "supergroup" {
		return nil
	}
	// If whitelisted groups are enabled
	if F.Bool("whitelist_mode") {
		// Stop serving non-whitelisted groups
		if WhiteList[c.Chat().ID] != 1 {
			// leave group
			return c.Bot().Leave(c.Chat())
		}
	}
	if c.Message().Sticker.SetName == "" {
		return c.Delete()
	}

	rule := StringMap(rd.Get(ctx, "sticker_Rule_"+cast.ToString(c.Chat().ID)).Result())
	modetype, _ := rd.Get(ctx, "sticker_Config_Mode_"+cast.ToString(c.Chat().ID)).Bool()

	// If whitelist mode is enabled
	if modetype {
		// delete it
		if rule[c.Message().Sticker.SetName] != "v" {
			return c.Delete()
		}
		// Skip it if present in the whitelist
		return nil
	}

	// Blacklist mode enabled

	// delete it
	if rule[c.Message().Sticker.SetName] == "v" {
		return c.Delete()
	}
	// Skip it if it is not in the blacklist
	return nil
}
