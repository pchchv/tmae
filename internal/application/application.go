package application

import (
	"io/fs"

	"github.com/go-telegram/bot"
)

type Application struct {
	botToken string
	subbedFS fs.FS
	b        *bot.Bot
}
