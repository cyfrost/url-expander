package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var mutex = &sync.Mutex{}

// PORT returns the port assigned by OS env vars
var PORT = os.Getenv("PORT")

// HOST returns the hostname returned by OS env vars
var HOST = os.Getenv("HOST")

func init() {
	if PORT == "" {
		PORT = "3000"
		// log.Fatalln("\nError: Env variable $PORT not set!")
	}

	if HOST == "" {
		HOST = "localhost"
		// log.Fatalln("\nError: Env variable $HOST not set!")
	}
}

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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, h *http.Request) {
		shortURL := h.URL.Query().Get("shorturl")

		// If the file doesn't exist, create it, or append to the file
		f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal(err)
		}

		appendData := time.Now().Format("2006-01-02 15:04:05") + " - " + shortURL + "\n"

		if _, err := f.Write([]byte(appendData)); err != nil {
			f.Close()
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}

		if shortURL != "" {
			mutex.Lock()
			res, err := ExpandURL(shortURL)
			mutex.Unlock()
			if err != nil {
				fmt.Fprintf(w, `{"error": "`+err.Error()+`"}`)
			} else {
				fmt.Fprintf(w, `{"result": "`+res+`"}`)
			}
		} else {
			fmt.Fprintf(w, `<html><head><title>URL Expander</title></head><body style="background: black; color: white;"><h3>This microservice is meant to be a slug against privacy deterrence. Short URLs are usually embedded little links that track the users click behavior, IP and more. This is a simple GoLang program that needs to be hosted on a VPS or something and then use the API to get an expanded link, this way the offending service will see the IP of the VPS instead of your own (thus protecting your privacy). You can also create a bot around this microservice and use it in many places (example: a Telegram bot). To get the expanded URL of any short URL, Send a GET request like this: http://localhost:< PORT_HERE >?shorturl=< SHORT_URL_HERE >. You'll receive a JSON in response with the expanded URL in it under the key "result" (or "error" if any).`)
		}
	})

	fmt.Println("Server started on http(s?)://" + HOST + ":" + PORT + "...")
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), router)
}
