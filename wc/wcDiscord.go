package wc

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = make(map[string]CommandHandler)

type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

func RegisterHandlers(handlers map[string]CommandHandler) {
	for name, handler := range handlers {
		CommandHandlers[name] = handler
	}
}

func botHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	fmt.Printf("Interaction received: Type: %v | ID: %v | User: %v#%v | Guild: %v\n",
		i.Type,
		i.ID,
		i.Member.User.Username,
		i.Member.User.Discriminator,
		i.GuildID,
	)
	if i.Type == discordgo.InteractionApplicationCommand {
		commandName := i.ApplicationCommandData().Name
		if handler, exists := CommandHandlers[commandName]; exists {
			handler(s, i)
		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Unknown command.",
				},
			})
		}
	}
}

func DgoDeferMsg(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource, // Deferring response
	})
	if err != nil {
		log.Println("Error deferring interaction:", err)
		return
	}
	log.Println("Interaction deferred successfully, starting processing...")
}

func DgoEmbedMsg(t string, v string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title: t,
		Color: 0x5797a3,
		Fields: []*discordgo.MessageEmbedField{
			{
				Value: v,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Powered by iWait Bot",
		},
	}

	_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{embed},
	})
	if err != nil {
		fmt.Println("Error editing response with embed:", err)
	}
}
