package wc

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type BotContext struct {
	Session     *discordgo.Session
	Interaction *discordgo.InteractionCreate
}

var CommandHandlers = make(map[string]func(ctx *BotContext))

type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

func RegisterHandlers(handlers map[string]func(ctx *BotContext)) {
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
			ctx := &BotContext{
				Session:     s,
				Interaction: i,
			}
			handler(ctx)
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

func (ctx *BotContext) DgoDeferMsg() {
	err := ctx.Session.InteractionRespond(ctx.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		log.Println("Error deferring interaction:", err)
		return
	}
	log.Println("Interaction deferred successfully, starting processing...")
}

func (ctx *BotContext) DgoEmbedMsg(title, value, footer string, options ...bool) {
	ephemeral := false
	if len(options) > 0 {
		ephemeral = options[0]
	}

	embed := &discordgo.MessageEmbed{
		Title: title,
		Color: 0x5797a3,
		Fields: []*discordgo.MessageEmbedField{
			{
				Value: value,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: footer,
		},
	}

	if ephemeral {
		err := ctx.Session.InteractionRespond(ctx.Interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
				Flags:  discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			fmt.Println("Error responding with ephemeral embed:", err)
		}
	} else {
		_, err := ctx.Session.InteractionResponseEdit(ctx.Interaction.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embed},
		})
		if err != nil {
			fmt.Println("Error editing response with embed:", err)
		}
	}
}
