package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"sync"
	"time"
)

var (
	bot  *tb.Bot
	once sync.Once
)

func Bot() *tb.Bot {
	once.Do(CreateBot)
	return bot
}

func CreateBot() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("BOT_KEY"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	bot = b
}

func SendVideoByURL(URL string, caption string, message *tb.Message) {
	video := &tb.Video{
		File:    tb.FromURL(URL),
		Caption: caption,
	}

	_, err := Bot().Send(message.Sender, video)
	if err != nil {
		log.Print(video)
		log.Print(err)
	}
}

func SendPictureByURL(URL string, caption string, message *tb.Message) {
	photo := &tb.Photo{
		File:    tb.FromURL(URL),
		Caption: caption,
	}

	_, err := Bot().Send(message.Sender, photo)
	if err != nil {
		log.Print(photo)
		log.Print(err)
	}
}

func ReportError(error error, message *tb.Message) {
	_, err := Bot().Send(message.Sender, "A fehler is passiert bitte schau da in Log au!")
	log.Print(error)
	if err != nil {
		log.Print(err)
	}
}

func SendText(text string, message *tb.Message) {
	_, err := Bot().Send(message.Sender, text)
	if err != nil {
		log.Print(err)
	}
}
