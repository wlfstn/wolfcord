package main

import (
	"flag"
	"fmt"

	"github.com/wlfstn/wolfcord/wc"
)

func main() {
	configPath := flag.String("config", "./resources/config.toml", "Path to the configuration file")
	commandsDir := flag.String("commands", "resources/commands", "Path to the commands directory")
	flag.Parse()

	wc.RegisterHandlers(map[string]func(ctx *wc.BotContext){
		"ping": handlePing,
	})
	wc.InitializeBot(*configPath, *commandsDir)
}

func handlePing(ctx *wc.BotContext) {
	fmt.Println("Running Ping")
	userName := ctx.Interaction.Member.User.Username

	title := "You Pinged?"
	value := fmt.Sprintf("Well Pong! Hello, %s!", userName)
	footer := "Powered by WolfCord Framework"
	ephemeral := true

	ctx.DgoEmbedMsg(title, value, footer, ephemeral)
}
