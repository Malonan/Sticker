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

	"sticker/lib/libg/dbstr"
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

	t := tb.New().SetContext(c)

	if err := packet1(t, 1); err != nil {
		return nil
	}

	t.SetAutoDelete(10).Send("Refreshing admin list...")
	if err := GetAdminList(c); err != nil {
		t.SetAutoDelete(12).Send(err.Error())
	}
	return nil
}

func CommandSelectMode(c tele.Context) error {
	t := tb.New().SetContext(c)
	if err := packet1(t); err != nil {
		return nil
	}
	c.Delete()
	modetype, _ := rd.Get(ctx, "sticker_Config_Mode_"+cast.ToString(c.Chat().ID)).Bool()
	if modetype {
		rd.Set(ctx, "sticker_Config_Mode_"+cast.ToString(c.Chat().ID), false, 0)
		db.Select("Modetype").Updates(&dbstr.Config{Gid: c.Chat().ID, Modetype: false})
		t.SetAutoDelete(12).Send("Group sticker checking mode has been switched to blacklist mode!")
		return nil
	}
	rd.Set(ctx, "sticker_Config_Mode_"+cast.ToString(c.Chat().ID), true, 0)
	db.Select("Modetype").Updates(&dbstr.Config{Gid: c.Chat().ID, Modetype: true})
	t.SetAutoDelete(12).Send("Group sticker checking mode has been switched to whitelist mode!")
	return nil
}