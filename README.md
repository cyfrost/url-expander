# URL Expander

### What
A Go microservice API to unshorten URLs

### Why
Short URLs are often very detrimental to user privacy in that they do click tracking (mapped to users' IP) and more. This microservice is meant to be a slug against that. This is a simple GoLang API that needs to be hosted on a VPS/server of some sort and then you can supply the shorturl to the API and get the expanded link as a result, this way the offending (URL shortener) service will see the IP of the VPS instead of your own (thus protecting your privacy). You can also create a bot around this microservice and use it in many places (example: a Telegram bot --has wrapper for Go). To get the expanded URL of any short URL, Send a GET request like this: < scheme > < host > :< port >?shorturl=< SHORT_URL_HERE >. You'll receive a JSON in response with the expanded URL in it under the key "result" (or "error" if any).

### How
1. install Go, and Git on your machine
2. clone this repository: `git clone https://github.com/cyfrost/url-expander`
3. pull deps: `go install url-expander.go`
4. run: `go run url-expander.go`

Once you mix-and-match your setup, try setting up a VPS with this and optionally build (bots?) around this microservice.

## Credits

1. URL expansion code used from https://github.com/eldadru/url-expander (thanks @eldadru)
2. http server code used from https://github.com/ishanjain28/url-shortner (thanks @ishanjain28)
