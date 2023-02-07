package src

/*
  Copyright 2023 Malonan & 3JoB

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

import (
	tele "github.com/3JoB/telebot"
	tb "github.com/3JoB/ulib/telebot/utils"
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

func ServiceJoinToGroup(c tele.Context) error {
	if c.Chat().Type != "supergroup" {
		return nil
	}
	t := tb.New().SetContext(c)
	msg := `Sticker admins have joined the group, are you ready to be filled with my wrath? (just kidding)
You can deploy an identical instance at https://github.com/Malonan/Sticker, or you can learn how to use this bot.`
	t.Send(msg)
	Add(c.Chat().ID)
	return GetAdminList(c)
}
