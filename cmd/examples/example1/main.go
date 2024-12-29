package main

import (
	"flag"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/wlfstn/wolfcord/wc"
)

func main() {
	configPath := flag.String("config", "./resources/config.toml", "Path to the configuration file")
	commandsDir := flag.String("commands", "resources/commands", "Path to the commands directory")
	flag.Parse() // Parse the command-line flags

	wc.RegisterHandlers(map[string]wc.CommandHandler{
		"ping": handlePing,
	})
	wc.InitializeBot(*configPath, *commandsDir)
}

func handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	fmt.Println("Running Ping")
	userName := i.Member.User.Username

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Pong! %s\n", userName),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
