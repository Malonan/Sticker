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
	tele "github.com/3JoB/telebot"
	"github.com/spf13/cast"

	"sticker/lib/libg/dbstr"
)

func Add(g int64) {
	db.Create(&dbstr.Config{Gid: g, Data: "{}", Admin: "{}", Modetype: false})
}

func CommandStickerBan(c tele.Context) error {
	c.Delete()
	// Check if the chat is a supergroup
	if c.Chat().Type != "supergroup" {
		return fn.SA(c, 12, "This command can only be used within a supergroup!!!")
	}
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
	// Must reply to a message with command
	if !c.Message().IsReply() {
		return fn.SA(c, 10, "Please use this command to reply to a message!!!")
	}
	// If the message object is not a sticker
	if c.Message().ReplyTo.Sticker == nil {
		return fn.SA(c, 12, "The selected object must be a sticker!!!")
	}
	// If the message object is not a sticker
	if c.Message().ReplyTo.Sticker.SetName == "" {
		return fn.SA(c, 12, "The selected object must be a sticker!!!")
	}

	// I hate the map[string]string bug in 1.20.
	// It makes the following code completely impossible to run!!!

	rule := StringMap(rd.Get(ctx, "sticker_Rule_"+cast.ToString(c.Chat().ID)).Result())
	if rule[c.Message().ReplyTo.Sticker.SetName] == "" {
		rule[c.Message().ReplyTo.Sticker.SetName] = "v"
		rules := String(rule)
		rd.Set(ctx, "sticker_Rule_"+cast.ToString(c.Chat().ID), rules, 0)
		db.Updates(&dbstr.Config{Gid: c.Chat().ID, Data: rules})
		return fn.SA(c, 8, "I already remember you!!!")
	}
	delete(rule, c.Message().ReplyTo.Sticker.SetName)
	rules := String(rule)
	db.Updates(&dbstr.Config{Gid: c.Chat().ID, Data: rules})
	rd.Set(ctx, "sticker_Rule_"+cast.ToString(c.Chat().ID), rules, 0)
	return fn.SA(c, 12, "my memory of it has faded...")
}
