package wc

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Config struct {
	KeyLoc    string
	ServerID  string
	ChannelID string
	Database  DatabaseConfig
}

var conf Config
var AuthChannels map[string]string
var AuthUsers map[string]string

type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

var (
	CmdsList []*discordgo.ApplicationCommand
	err      error
)

func initBot() {
	CmdsList, err = LoadCmdsFromTOML("./resources/commands")
	if err != nil {
		fmt.Println("Error loading commands:", err)
	}
}

func verifySlashCommands(dg *discordgo.Session) {
	// Create a map of existing commands presently stored in Discord
	log.Println("Verifying commands...")

	mapExistingCommands := make(map[string]*discordgo.ApplicationCommand)
	existingCmds, err := dg.ApplicationCommands(dg.State.User.ID, conf.ServerID)
	if err != nil {
		log.Println("Error retrieving commands: ", err)
	}

	for _, cmd := range existingCmds {
		mapExistingCommands[cmd.Name] = cmd
	}

	mapNewCommands := make(map[string]*discordgo.ApplicationCommand)
	for _, cmd := range CmdsList {
		mapNewCommands[cmd.Name] = cmd
	}

	diff := MapUpdateCompare(mapNewCommands, mapExistingCommands)
	for cmdName, cmdStatus := range diff {
		switch cmdStatus {
		case Added:
			log.Printf("Command %s does not exist, creating...\n", cmdName)
			_, err := dg.ApplicationCommandCreate(dg.State.User.ID, conf.ServerID, mapNewCommands[cmdName])
			if err != nil {
				log.Println("Error creating command: ", err)
			}
		case Removed:
			log.Printf("Command %s no longer exists, removing...\n", cmdName)
			err := dg.ApplicationCommandDelete(dg.State.User.ID, conf.ServerID, mapExistingCommands[cmdName].ID)
			if err != nil {
				log.Println("Error deleting command: ", err)
			}
		case Updated:
			log.Printf("Command %s has changed, updating...\n", cmdName)
			_, err := dg.ApplicationCommandEdit(dg.State.User.ID, conf.ServerID, mapExistingCommands[cmdName].ID, mapNewCommands[cmdName])
			if err != nil {
				log.Println("Error updating command: ", err)

			}
		case Equal:
			log.Printf("Command %s is up to date.\n", cmdName)
		}
	}
}
