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
	"github.com/tidwall/gjson"

	"sticker/lib/libg/dbstr"
)

type AdminList struct {
	User struct {
		ID int64 `json:"id"`
	} `json:"user"`
}

func Int64Map(a string, x ...any) map[int64]int {
	smp := make(map[int64]int)
	json.UnmarshalString(a, &smp)
	return smp
}

func StringMap(a string, x ...any) map[string]string {
	smp := make(map[string]string)
	json.UnmarshalString(a, &smp)
	return smp
}

// Send files with tele.Bot
/*func (*Func) FS(c tele.Context, u int64, msg any) (*tele.Message, error) {
	return c.Bot().Send(tele.ChatID(u), msg)
}

// Send messages with tele.Bot
func (*Func) CS(c tele.Context, u int64, msg ...any) (*tele.Message, error) {
	if len(msg) == 2 {
		return c.Bot().Send(tele.ChatID(u), msg[0], &tele.SendOptions{ParseMode: tele.ModeHTML, DisableWebPagePreview: true}, msg[1])
	}
	return c.Bot().Send(tele.ChatID(u), msg[0], &tele.SendOptions{ParseMode: tele.ModeHTML, DisableWebPagePreview: true})
}

// Edit Message
func (*Func) Edit(c tele.Context, msg ...any) error {
	if len(msg) == 2 {
		return c.Edit(msg[0], &tele.SendOptions{ParseMode: tele.ModeHTML, DisableWebPagePreview: true}, msg[1])
	}
	return c.Edit(msg[0], &tele.SendOptions{ParseMode: tele.ModeHTML, DisableWebPagePreview: true})
}*/

// Get the list of group administrators
func GetAdminList(c tele.Context) error {
	var b []AdminList
	m := map[string]int64{
		"chat_id": c.Chat().ID,
	}
	t := tb.New().SetContext(c)
	d, _ := c.Bot().Raw("getChatAdministrators", m)
	if !gjson.GetBytes(d, "ok").Bool() {
		return nil
	}
	json.UnmarshalString(gjson.GetBytes(d, "result").String(), &b)
	if len(b) == 0 {
		t.SetAutoDelete(10).Send("Oh no....bot failed to fetch admin list....please check what happened....")
		return nil
	}
	admin := make(map[int64]int)
	for _, i := range b {
		admin[i.User.ID] = 1
	}
	rd.Set(ctx, "sticker_Admin_"+cast.ToString(c.Chat().ID), json.Marshal(&admin).String(), 0)
	db.Updates(&dbstr.Config{Gid: c.Chat().ID, Admin: json.Marshal(&admin).String()})
	t.SetAutoDelete(10).Send("The admin list has been refreshed.")
	return nil
}

// Format Timestamp
/*func (*Func) Format(t int64) (date string) {
	obj := time.Unix(int64(t), 0)
	date = obj.Format("2006-01-02 15:04:05 GMT+0")
	return
}
*/
