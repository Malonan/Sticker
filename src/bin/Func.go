package bin

import (
	"time"

	tele "github.com/3JoB/telebot"
	"github.com/goccy/go-json"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"

	"sticker/lib/libg/dbstr"
)

type Func struct{}

func String(a any) string {
	s, err := json.Marshal(&a)
	if err != nil {
		return ""
	}
	return cast.ToString(s)
}

func Int64Map(a string, x ...any) map[int64]int {
	smp := make(map[int64]int)
	json.Unmarshal([]byte(a), &smp)
	return smp
}

func StringMap(a string, x ...any) map[string]string {
	smp := make(map[string]string)
	json.Unmarshal([]byte(a), &smp)
	return smp
}

// Delete Message
func (*Func) D(c tele.Context, message int, chat int64) error {
	return c.Bot().Delete(&tele.StoredMessage{
		MessageID: cast.ToString(message),
		ChatID:    chat,
	})
}

// Send Message
func (*Func) S(c tele.Context, msg ...any) (*tele.Message, error) {
	if len(msg) == 2 {
		return c.Send(msg[0], &tele.SendOptions{ParseMode: tele.ModeHTML, DisableWebPagePreview: true}, msg[1])
	}
	return c.Send(msg[0], &tele.SendOptions{ParseMode: tele.ModeHTML, DisableWebPagePreview: true})
}

// Send an auto-destroy message
func (*Func) SA(c tele.Context, deletetime int, msg any) error {
	i, _ := fn.S(c, msg)
	time.Sleep(time.Second * time.Duration(deletetime))
	fn.D(c, i.ID, c.Chat().ID)
	return nil
}

// Send files with tele.Bot
func (*Func) FS(c tele.Context, u int64, msg any) (*tele.Message, error) {
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
}

//Pop-ups
func (*Func) Alert(c tele.Context, text string) error {
	return c.Respond(&tele.CallbackResponse{
		Text:      text,
		ShowAlert: true,
	})
}

//Prompt information
func (*Func) AlertNo(c tele.Context, text string) error {
	return c.Respond(&tele.CallbackResponse{
		Text:      text,
		ShowAlert: false,
	})
}

//Get the list of group administrators
func GetAdminList(c tele.Context) error {
	var b []AdminList
	m := map[string]int64{
		"chat_id": c.Chat().ID,
	}
	d, _ := c.Bot().Raw("getChatAdministrators", m)
	if !gjson.GetBytes(d, "ok").Bool() {
		return nil
	}
	json.Unmarshal([]byte(gjson.GetBytes(d, "result").String()), &b)
	if len(b) == 0 {
		fn.SA(c, 10,"Oh no....bot failed to fetch admin list....please check what happened....")
		return nil
	}
	admin := make(map[int64]int)
	for _, i := range b {
		admin[i.User.ID] = 1
	}
	rd.Set(ctx, "sticker_Admin_"+cast.ToString(c.Chat().ID), String(admin), 0)
	db.Updates(&dbstr.Config{Gid: c.Chat().ID, Admin: String(admin)})
	fn.SA(c, 10, "The admin list has been refreshed.")
	return nil
}

// Format Timestamp
func (*Func) Format(t int64) (date string) {
	obj := time.Unix(int64(t), 0)
	date = obj.Format("2006-01-02 15:04:05 GMT+0")
	return
}
