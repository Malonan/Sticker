package bin

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
	"sticker/lib/libg/dbstr"

	tele "github.com/3JoB/telebot"
	"github.com/spf13/cast"
)

func CommandStart(c tele.Context) error {
	if c.Chat().Type != "private" {
		return nil
	}
	msg := `This is me! Are you ready to be filled with my wrath? (just kidding)
You can deploy an identical instance at https://github.com/Malonan/Sticker, or you can learn how to use this bot.`
	c.Send(msg)
	return nil
}

func CommandRefresh(c tele.Context) error {
	if c.Chat().Type != "supergroup" {
		c.Send("What?")
		return nil
	}
	c.Delete()
	// If whitelisted groups are enabled
	if F.Bool("whitelist_mode") {
		// Stop serving non-whitelisted groups
		if WhiteList[c.Chat().ID] != 1 {
			fn.S(c, "This group is not available for this function!!!")
			// leave group
			return c.Bot().Leave(c.Chat())
		}
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

func CommandSelectMode(c tele.Context) error {
	// Check if the chat is a supergroup
	if c.Chat().Type != "supergroup" {
		return fn.SA(c, 12, "This command can only be used within a supergroup!!!")
	}
	c.Delete()
	// If whitelisted groups are enabled
	if F.Bool("whitelist_mode") {
		// Stop serving non-whitelisted groups
		if WhiteList[c.Chat().ID] != 1 {
			fn.S(c, "This group is not available for this function!!!")
			// leave group
			return c.Bot().Leave(c.Chat())
		}
	}
	admin := Int64Map(rd.Get(ctx, "sticker_Admin_"+cast.ToString(c.Chat().ID)).Result())
	// Prevent non-admins from operating the bot
	if admin[c.Sender().ID] != 1 {
		return fn.SA(c, 10, "This command is only available to supergroup administrators!!!")
	}
	modetype , _ := rd.Get(ctx, "sticker_Config_Mode_"+cast.ToString(c.Chat().ID)).Bool()
	if modetype {
		rd.Set(ctx, "sticker_Config_Mode_"+cast.ToString(c.Chat().ID), false, 0)
		db.Updates(&dbstr.Config{Gid: c.Chat().ID, Modetype: false})
		return fn.SA(c, 12, "Group sticker checking mode has been switched to blacklist mode!")
	}
		rd.Set(ctx, "sticker_Config_Mode_"+cast.ToString(c.Chat().ID), true, 0)
		db.Updates(&dbstr.Config{Gid: c.Chat().ID, Modetype: true})
		return fn.SA(c, 12, "Group sticker checking mode has been switched to whitelist mode!")
}