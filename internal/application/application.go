package application

import (
	"embed"
	"errors"
	"io/fs"
	"log"
	"net/http"

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

func (app *Application) handlerAPIOpen(rw http.ResponseWriter, req *http.Request) {
	if user, ok := bot.ValidateWebappRequest(req.URL.Query(), app.botToken); !ok {
		http.Error(rw, "unauthorized", http.StatusUnauthorized)
	} else {
		log.Printf("%v", user)
	}
}
