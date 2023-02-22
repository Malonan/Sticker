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
	"github.com/3JoB/ulib/json"
	tb "github.com/3JoB/ulib/telebot/utils"
	"github.com/spf13/cast"

	"sticker/lib/db/dbstr"
)

func Add(g int64) {
	dbs.Create(&dbstr.Config{Gid: g, Data: "{}", Admin: "{}", Modetype: false})
}

func CommandStickerBan(c tele.Context) error {
	// fmt.Println(json.Marshal(c.Message()).String())
	t := tb.New().SetContext(c)
	if c.Chat().IsForum {
		t = t.SetTopicID(int64(c.Message().ThreadID))
	}
	if err := packet(tb.New().SetContext(c)); err == errs {
		return nil
	}
	// Must reply to a message with command
	if !c.Message().IsReply() {
		CheckErr(t.SetAutoDelete(10).SetDeleteCommand().Send("Please use this command to reply to a message!!!"))
		return nil
	}
	// If the message object is not a sticker
	if c.Message().ReplyTo.Sticker == nil {
		CheckErr(t.SetAutoDelete(12).SetDeleteCommand().Send("The selected object must be a sticker!!!"))
		return nil
	}
	// If the message object is not a sticker
	if c.Message().ReplyTo.Sticker.SetName == "" {
		CheckErr(t.SetAutoDelete(12).SetDeleteCommand().Send("The selected object must be a sticker!!!"))
		return nil
	}

	// I hate the map[string]string bug in 1.20.
	// It makes the following code completely impossible to run!!!

	t = tb.New().SetContext(c)

	rule := RuleMap(c.Chat().ID)
	if rule[c.Message().ReplyTo.Sticker.SetName] == "" {
		rule[c.Message().ReplyTo.Sticker.SetName] = "v"
		rules := json.Marshal(&rule).String()
		rd.Set(ctx, "sticker_Rule_"+cast.ToString(c.Chat().ID), rules, 0)
		dbs.Updates(&dbstr.Config{Gid: c.Chat().ID, Data: rules})
		CheckErr(t.SetAutoDelete(10).Send("I already remember you!!!"))
		return nil
	}
	delete(rule, c.Message().ReplyTo.Sticker.SetName)
	rules := json.Marshal(&rule).String()
	dbs.Updates(&dbstr.Config{Gid: c.Chat().ID, Data: rules})
	rd.Set(ctx, "sticker_Rule_"+cast.ToString(c.Chat().ID), rules, 0)
	CheckErr(t.SetAutoDelete(10).Send("my memory of it has faded..."))
	return nil
}
