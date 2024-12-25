package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/wlfstn/wolfcord/wc"
)

func main() {
	configLoc := "./resources/config.toml"
	commandsDir := "./resources/commands"

	wc.RegisterHandlers(map[string]wc.CommandHandler{
		"ping": handlePing,
	})
	wc.InitializeBot(configLoc, commandsDir)
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
