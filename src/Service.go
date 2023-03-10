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
	if err := CheckPerm(c); err != nil {
		if err == errs {
			return nil
		}
	}

	if c.Message().Sticker.SetName == "" {
		return c.Delete()
	}

	rule := RuleMap(c.Chat().ID)
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
	if c.Chat().Type != tele.ChatSuperGroup || c.Chat().Type != tele.ChatGroup {
		return nil
	}
	t := tb.New().SetContext(c)
	msg := `Sticker admins have joined the group, are you ready to be filled with my wrath? (just kidding)
You can deploy an identical instance at https://github.com/Malonan/Sticker, or you can learn how to use this bot.`
	t.Send(msg)
	Add(c.Chat().ID)
	if err := GetAdminList(c); err != nil {
		t.SetAutoDelete(10).Send(err.Error())
	}
	return nil
}
