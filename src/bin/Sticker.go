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
	"github.com/3JoB/ulib/json"
	tb "github.com/3JoB/ulib/telebot"
	"github.com/spf13/cast"

	"sticker/lib/libg/dbstr"
)

func Add(g int64) {
	db.Create(&dbstr.Config{Gid: g, Data: "{}", Admin: "{}", Modetype: false})
}

func CommandStickerBan(c tele.Context) error {
	c.Delete()
	t := tb.New().SetContext(c)
	// Check if the chat is a supergroup
	if c.Chat().Type != "supergroup" {
		t.SetAutoDelete(12).Send("This command can only be used within a supergroup!!!")
		return nil
	}
	// If whitelisted groups are enabled
	if F.Bool("whitelist_mode") {
		// Stop serving non-whitelisted groups
		if WhiteList[c.Chat().ID] != 1 {
			t.Send("This group is not available for this function!!!")
			// leave group
			return c.Bot().Leave(c.Chat())
		}
	}

	admin := Int64Map(rd.Get(ctx, "sticker_Admin_"+cast.ToString(c.Chat().ID)).Result())
	// Prevent non-admins from operating the bot
	if admin[c.Sender().ID] != 1 {
		t.SetAutoDelete(10).Send("This command is only available to supergroup administrators!!!")
		return nil
	}
	// Must reply to a message with command
	if !c.Message().IsReply() {
		t.SetAutoDelete(10).Send("Please use this command to reply to a message!!!")
		return nil
	}
	// If the message object is not a sticker
	if c.Message().ReplyTo.Sticker == nil {
		t.SetAutoDelete(12).Send("The selected object must be a sticker!!!")
		return nil
	}
	// If the message object is not a sticker
	if c.Message().ReplyTo.Sticker.SetName == "" {
		t.SetAutoDelete(12).Send("The selected object must be a sticker!!!")
		return nil
	}

	// I hate the map[string]string bug in 1.20.
	// It makes the following code completely impossible to run!!!

	rule := StringMap(rd.Get(ctx, "sticker_Rule_"+cast.ToString(c.Chat().ID)).Result())
	if rule[c.Message().ReplyTo.Sticker.SetName] == "" {
		rule[c.Message().ReplyTo.Sticker.SetName] = "v"
		rules := json.Marshal(&rule).String()
		rd.Set(ctx, "sticker_Rule_"+cast.ToString(c.Chat().ID), rules, 0)
		db.Updates(&dbstr.Config{Gid: c.Chat().ID, Data: rules})
		t.SetAutoDelete(10).Send("I already remember you!!!")
		return nil
	}
	delete(rule, c.Message().ReplyTo.Sticker.SetName)
	rules := json.Marshal(&rule).String()
	db.Updates(&dbstr.Config{Gid: c.Chat().ID, Data: rules})
	rd.Set(ctx, "sticker_Rule_"+cast.ToString(c.Chat().ID), rules, 0)
	t.SetAutoDelete(10).Send("my memory of it has faded...")
	return nil
}
