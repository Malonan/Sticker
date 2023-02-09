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
	"errors"

	tele "github.com/3JoB/telebot"
	"github.com/3JoB/ulib/json"
	tb "github.com/3JoB/ulib/telebot/utils"
	"github.com/spf13/cast"

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

type AdminRule map[int64]tb.AdminInfo

func AdminMap(a string, x ...any) AdminRule{
	info := make(AdminRule)
	json.UnmarshalString(a, &info)
	return info
}

var (
	errs = errors.New("pe")
)

func packet1(t *tb.Use, ref ...any) error {
	// Check if the chat is a supergroup
	if t.Ctx.Chat().Type != "supergroup" {
		t.SetAutoDelete(12).Send("This command can only be used within a supergroup!!!")
		return errs
	}
	// If whitelisted groups are enabled
	if F.Bool("whitelist_mode") {
		// Stop serving non-whitelisted groups
		if WhiteList[t.Ctx.Chat().ID] != 1 {
			t.Send("This group is not available for this function!!!")
			// leave group
			t.Ctx.Bot().Leave(t.Ctx.Chat())
			return errs
		}
	}

	admin := AdminMap(rd.Get(ctx, "sticker_Admin_"+cast.ToString(t.Ctx.Chat().ID)).Result())
	if len(ref) != 0 {
		if len(admin) == 0 {
			if err := GetAdminList(t.Ctx); err != nil {
				t.SetAutoDelete(10).Send(err.Error())
				return errs
			}
			t.SetAutoDelete(10).Send("The current group management list is empty and is trying to get it.\nIf you can't get it, check that the bot has been granted administrator privileges.")
			return errs
		}
	}

	// Prevent non-admins from operating the bot
	if admin[t.Ctx.Bot().Me.ID].User.ID == 0 {
		t.SetAutoDelete(10).Send("The robot is not a group administrator, the operation is not available.")
		return errs
	}
	if !admin[t.Ctx.Bot().Me.ID].CanDeleteMessages {
		t.SetAutoDelete(10).Send("Insufficient permissions for the robot to operate.")
		return errs
	}
	// Prevent non-admins from operating the bot
	if admin[t.Ctx.Sender().ID].User.ID == 0 {
		t.SetAutoDelete(10).Send("This command is only available to supergroup administrators!!!")
		return errs
	}
	return nil
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
	t, err := tb.New().SetContext(c).SetChatID(c.Chat().ID).GetAdminList()
	if err != nil {
		return err
	}
	rd.Set(ctx, "sticker_Admin_"+cast.ToString(c.Chat().ID), json.Marshal(&t).String(), 0)
	db.Updates(&dbstr.Config{Gid: c.Chat().ID, Admin: json.Marshal(&t).String()})
	tb.New().SetContext(c).SetAutoDelete(10).Send("The admin list has been refreshed.")
	return nil
}

// Format Timestamp
/*func (*Func) Format(t int64) (date string) {
	obj := time.Unix(int64(t), 0)
	date = obj.Format("2006-01-02 15:04:05 GMT+0")
	return
}
*/
