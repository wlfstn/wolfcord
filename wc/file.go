package wc

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string          `toml:"name"`
	Description string          `toml:"description"`
	Options     []CommandOption `toml:"options"`
}

type CommandOption struct {
	Name         string          `toml:"name"`
	Description  string          `toml:"description"`
	Type         int             `toml:"type"`
	Required     bool            `toml:"required"`
	Autocomplete bool            `toml:"autocomplete"`
	Options      []CommandOption `toml:"options"`
}

func ReadFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("file error: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("read error: %v", err)
	}

	output := strings.TrimSpace(string(content))
	return output
}

func ReadQueryFromFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	defer file.Close()

	query, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(query)
}

func LoadCmdsFromTOML(dir string) ([]*discordgo.ApplicationCommand, error) {
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
