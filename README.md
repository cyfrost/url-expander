# URL Expander

### What
A tiny Go microservice API to unshorten URLs

### Why
Short URLs are often detrimental to user privacy in that they do click tracking (mapped to users' pub IP) and more. This microservice is meant to be a slug against that. This is a simple GoLang API that needs to be hosted on a VPS/server of some sort and then you can supply the shorturl to the API and get the expanded link as a result, this way the offending (URL shortener) service will see the IP of the VPS instead of your own (thus protecting your privacy). You can also create a bot around this microservice and use it in many places (example: a Telegram bot --has wrapper for Go). To get the expanded URL of any short URL, Send a GET request like this: < scheme > < host > :< port >?shorturl=< SHORT_URL_HERE >. You'll receive a JSON in response with the expanded URL in it under the key "result" (or "error" if any).

### Usage Instructions

## If you want to run this as a Microservice

1. install Go language, and Git on your machine
2. clone this repository: `git clone https://github.com/cyfrost/url-expander`
3. pull deps: `go install local_server.go`
4. run: `go run local_server.go`
5. (optional) you can build it into a binary with `go build local_server.go` (For multi-arch/cross-compilation supply GOARCH flags alongside).

## If you want to run this as a Telegram Bot

1. Use the Telegram's BotFather bot to create a new bot, acquire a token (make it webhook based).
2. Replace the BotFather issued token in the `tg-bot.go` file.
3. Compile the program with `go build tg-bot.go`
4. Start the bot server on your VPS/any machine like this `./tg-bot`
5. Open Telegram and test if the server is responding to your messages as expected.

## Credits

1. URL expansion code used from https://github.com/eldadru/url-expander (thanks @eldadru)
2. http server code used from https://github.com/ishanjain28/url-shortner (thanks @ishanjain28)


## License

Apache v2
