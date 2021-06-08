simple Telegram Bot Api wrapper, maybe will grow to framework 

<h1>
This package is on alpha version,
and any update can broke backward capability
</h1>

NOTE: Please dont try use this package 

## version 0.1.1
<br>

## docs
<br>
 (WIP) for first time, you can read the code 

Download This package using this command `go get -v github.com/pikoUsername/tgp` 
Ignore a warning about no files in root directory, if you want download a new version of this package 
Then delete a past version in folder where saved on.

## Example
```go
package main

import (
	"fmt"
	"log"

	tgbot "github.com/pikoUsername/tgp/bot"
	"github.com/pikoUsername/tgp/dispatcher"
	"github.com/pikoUsername/tgp/utils"
    "github.com/pikoUsername/tgp/objects"
    "github.com/pikoUsername/tgp/configs"
)

func main() {
	bot, err := tgbot.NewBot("<token>", true, utils.ModeHTML)
	if err != nil {
		panic(err)
	}

	dp := dispatcher.NewDispatcher(bot)
	if err != nil {
		panic(err)
	}
    dp.MessageHandler.Register(func(u *objects.Update) { 
        if u.Message.Text == "" { 
            return 
        }
        _, err := bot.SendMessage(&configs.SendMessageConfig{
            ChatID: u.Message.Chat.ID, 
            Text: u.Message.Text, 
        })
        if err != nil { 
            fmt.Println(err)
        }
    })
    dp.StartPolling(dispatcher.NewStartPollingConfig(true))
}
```
