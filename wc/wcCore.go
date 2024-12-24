package wc

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
)

var (
	conf      Config
	AuthUsers map[string]string
	CmdsList  []*discordgo.ApplicationCommand
	err       error
)

// example dirPath "./resources/commmands"
func InitializeBot(tomlConfigs []string) {
	collectAuthUsers(&AuthUsers, tomlConfigs[1])

	_, err := toml.DecodeFile("config.toml", &conf)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("Server ID: %s\nChannel ID: %s\n", conf.ServerID, conf.ChannelID)

	postgresPass := ResourceLoadFile(conf.Database.Password)

	DbConn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Database.User, postgresPass, conf.Database.Host, conf.Database.Port, conf.Database.DBName)

	InitPGX(DbConn)

	discordToken := ResourceLoadFile(conf.KeyLoc)

	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	CmdsList, err = resourceLoadCommandFiles(tomlConfigs[2])
	if err != nil {
		fmt.Println("Error loading commands:", err)
	}

	verifySlashCommands(dg)
	fmt.Println("Bot is starting up!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Bot is shutting down!")
	dg.Close()
}

func collectAuthUsers(wlMap *map[string]string, file string) {
	var wl AuthUsersList

	_, err := toml.DecodeFile(file, &wl)
	if err != nil {
		fmt.Println("Whitelist Error:", err)
		os.Exit(1)
	}

	for _, entity := range wl.Entities {
		(*wlMap)[entity.ID] = entity.Name
	}
	fmt.Printf("Whitelist: %v\n", (*wlMap))
}

func resourceLoadCommandFiles(dir string) ([]*discordgo.ApplicationCommand, error) {
	var cmds []*discordgo.ApplicationCommand

	// Walk through the directory to find all TOML files
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		log.Println("Loading commands from:", path)
		if err != nil {
			return err
		}
		if filepath.Ext(path) != ".toml" {
			return nil
		}

		var cmd Command
		if _, err := toml.DecodeFile(path, &cmd); err != nil {
			return err
		}

		discordCmd := &discordgo.ApplicationCommand{
			Name:        cmd.Name,
			Description: cmd.Description,
		}

		var convertOptions func([]CommandOption) []*discordgo.ApplicationCommandOption
		convertOptions = func(options []CommandOption) []*discordgo.ApplicationCommandOption {
			var discordOptions []*discordgo.ApplicationCommandOption
			for _, opt := range options {
				discordOpt := &discordgo.ApplicationCommandOption{
					Name:         opt.Name,
					Description:  opt.Description,
					Type:         discordgo.ApplicationCommandOptionType(opt.Type),
					Required:     opt.Required,
					Autocomplete: opt.Autocomplete,
				}
				if len(opt.Options) > 0 {
					discordOpt.Options = convertOptions(opt.Options)
				}
				discordOptions = append(discordOptions, discordOpt)
			}
			return discordOptions
		}

		discordCmd.Options = convertOptions(cmd.Options)
		cmds = append(cmds, discordCmd)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cmds, nil
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