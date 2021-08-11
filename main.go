package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main() {
	_, err := discordgo.New("Bot " + "authentication token")
	if err != nil {
		fmt.Println("got some error")
	}
}