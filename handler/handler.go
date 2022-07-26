package handler

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	// "log"
	"net/http"
	"os"
	"strings"
)

// strings.Contains(search, target)（searchの中にtargetが含まれてるか）
// strings.EqualFold(s1, s2)（s1はs2と等しいか）
// strings.Index(search, target)（searchの中にtargetが含まれている場合のindex）
// strings.LastIndex(search, target)（searchの中にtargetが含まれている場合の最後のindex）
// strings.Replace(search, old, new, n)（searchの中のoldをnewに置換する）
// strings.ReplaceAll(search, old, new)（searchの中のoldをnewに置換する）
// strings.Split(s, sep)（sをsepで分割する）
// strings.ToLower(s)（sを小文字にする）
// strings.ToUpper(s)（sを大文字にする）
// strings.Trim(s, cutset)（sの前後にcutsetがある場合は削除する）
// strings.TrimLeft(s, cutset)（sの前にcutsetがある場合は削除する）
// strings.TrimRight(s, cutset)（sの後にcutsetがある場合は削除する）
// strings.TrimSpace(s)（sの前後に空白がある場合は削除する）
// strings.Join(values, sep)（valuesをsepで連結する）
// strings.Repeat(s, n)（sをn回繰り返す）
// strings.ToTitle(s)（sをタイトルケースにする）

func resultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "成功してます。")
}

func LINEHandler(w http.ResponseWriter, r *http.Request) {
  err := godotenv.Load("./.env")
  if err != nil {
    fmt.Println("Error loading .env file")
  }

  bot, err := linebot.New(
    os.Getenv("CHANNEL_SECRET"),
    os.Getenv("TOKEN"),
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		// イベントがメッセージの受信だった場合
		fmt.Println("イベント  ：", event)
		fmt.Println("タイプ　  ：", event.Type)
		fmt.Println("名前　　  ：", event.Source.UserID)
		fmt.Println("グループID：", event.Source.GroupID)
		fmt.Println("ルームID　：")
		fmt.Println("日時　　  ：", event.Timestamp)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// メッセージがテキスト形式の場合
			case *linebot.TextMessage:
				fmt.Println("メッセージ：", message)
				if strings.EqualFold(message.Text, "おはよう") {
					// メッセージを送信
					go bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("おはよう")).Do()
				} else if strings.EqualFold(message.Text, "こんにちは") {
					// メッセージを送信
					go bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("こんにちは")).Do()
				} else if strings.EqualFold(message.Text, "こんばんは") {
					// メッセージを送信
					go bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("こんばんは")).Do()
				}
			// メッセージが位置情報の場合
			case *linebot.LocationMessage:
				fmt.Println("メッセージ：", message)
				go bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("緯度: %f 経度: %f", message.Latitude, message.Longitude))).Do()
			// メッセージが画像の場合
			case *linebot.ImageMessage:
				fmt.Println("メッセージ：", message)
				go bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("画像が送信されました")).Do()
			//メッセージが音声の場合
			case *linebot.AudioMessage:
				fmt.Println("メッセージ：", message)
				go bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("音声が送信されました")).Do()
			//メッセージが動画の場合
			case *linebot.VideoMessage:
				fmt.Println("メッセージ：", message)
				go bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("動画が送信されました")).Do()
			//メッセージがファイルの場合
			case *linebot.FileMessage:
				fmt.Println("メッセージ：", message)
				go bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ファイルが送信されました")).Do()
			// メッセージがステッカーの場合
			case *linebot.StickerMessage:
				fmt.Println("メッセージ：", message)
				go bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ステッカーが送信されました")).Do()
			}
		}
	}
}
