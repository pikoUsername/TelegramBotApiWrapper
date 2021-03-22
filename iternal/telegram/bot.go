package telegram

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/pikoUsername/TelegramBotApiWrapper/iternal/telegram/types"
	"github.com/pikoUsername/TelegramBotApiWrapper/iternal/telegram/utils"
)

// Bot ...
type Bot struct {
	Token string

	Me types.User `json:"-"`
}

// TelegramURL ...
var TelegramURL string = "https://api.telegram.org/"

// NewBot ...
func NewBot(token string, checkToken bool) *Bot {
	if checkToken {
		err := utils.CheckToken(token)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &Bot{
		Token: token,
	}
}

// MakeRequest to telegram servers
// and result parses
func MakeRequest(Method string, Token string) (*types.TelegramResponse, error) {
	resp, err := http.Post(TelegramURL+"/"+Token+"/"+Method, "application/json", &strings.Reader{})
	if err != nil {
		return nil, errors.New("Error on sending request")
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(resp)
	if err != nil {
		return nil, errors.New("Error on json decode")
	}
	return resp, nil
}

// GetMe ...
func (bot *Bot) GetMe() *types.User {
	return nil
}

// SendMessage ...
func (b *Bot) SendMessage(text string) (*types.Message, error) {
	resp, err := MakeRequest("SendMessage", b.Token)
	if err != nil {
		return &types.Message{}, nil
	}
	json.Unmarshal(resp.Body)
	return &types.Message{}, nil
}