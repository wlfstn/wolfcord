package wc

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

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
