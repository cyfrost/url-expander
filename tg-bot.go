package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"time"
	"sync"
	"net/url"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var mutex = &sync.Mutex{}

// ExpandURL returns an expanded form of the short URL
func ExpandURL(url string) (string, error) {
	expandedURL := url

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			expandedURL = req.URL.String()
			return nil
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return expandedURL, nil
}

func initTgBotHandler() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeToken")

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		var responseText string

		userRequestedURL := update.Message.Text

		// If the file doesn't exist, create it, or append to the file
		f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal(err)
		}

		appendData := time.Now().Format("2006-01-02 15:04:05") + " - [Tg Bot] " + userRequestedURL + "\n"

		if _, err := f.Write([]byte(appendData)); err != nil {
			f.Close()
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}

		u, urlParseError := url.ParseRequestURI(userRequestedURL)
		
		if urlParseError != nil {
			responseText = "Erm, I speak only URL tongue, none other. Just send me a short URL and I'll expand it for ya! Why you ask? to protect your privacy from evil click tracking via short URLs. I would do that for ya? ofcourse, we're friends! Here ya go: https://github.com/cyfrost/url-expander"
		} else {
			mutex.Lock()
			fmt.Println(u)
			res, err := ExpandURL(userRequestedURL)
			mutex.Unlock()
			if err != nil {
				responseText = err.Error()
			} else {
				responseText = res
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func main() {
	initTgBotHandler()
}
