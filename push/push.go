package push

import (
	"log"
	"os"

	"github.com/gregdel/pushover"
)

type (
	push struct {
		Text string
	}
)

func (p push) Push() (*pushover.Response, error) {
	app := pushover.New(os.Getenv("PUSH_TOKEN"))
	recipient := pushover.NewRecipient(os.Getenv("PUSH_USER"))
	message := pushover.NewMessage(p.Text)
	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Println(err)
		return &pushover.Response{Status: 0}, err
	}
	return response, nil
}

func New() *push {
	return &push{}
}
