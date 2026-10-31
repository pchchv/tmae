package application

import (
	"embed"
	"errors"
	"io/fs"

	"github.com/go-telegram/bot"
)

//go:embed static
var staticContent embed.FS

type Application struct {
	botToken string
	subbedFS fs.FS
	b        *bot.Bot
}

func New(botToken string) (*Application, error) {
	app := &Application{
		botToken: botToken,
	}

	subbedFS, errSubFS := fs.Sub(staticContent, "static")
	if errSubFS != nil {
		return nil, errors.New("failed to create sub fs: " + errSubFS.Error())
	}
	app.subbedFS = subbedFS

	b, errBot := bot.New(botToken)
	if errBot != nil {
		return nil, errors.New("failed to create bot: " + errBot.Error())
	}

	app.b = b
	return app, nil
}
