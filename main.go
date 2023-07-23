package main

import (
	"fmt"

	"github.com/v1kr4nt/Discord-Bot/bot"
	"github.com/v1kr4nt/Discord-Bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	bot.Start()

	<-make(chan struct{})
	return
}
