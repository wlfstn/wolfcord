package wc

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const (
	Equal uint8 = iota
	Added
	Removed
	Updated
)

// Compare each element of two maps, return an array of differences
func MapUpdateCompare(newMap, oldMap map[string]*discordgo.ApplicationCommand) map[string]uint8 {
	diff := make(map[string]uint8)
	var modCheck []string

	for cmdName := range oldMap {
		if _, ok := newMap[cmdName]; !ok {
			diff[cmdName] = Removed
		} else {
			modCheck = append(modCheck, cmdName)
		}
	}

	for cmdName := range newMap {
		if _, ok := oldMap[cmdName]; !ok {
			diff[cmdName] = Added
		} else {
			modCheck = append(modCheck, cmdName)
		}
	}

	for _, cmdName := range modCheck {
		if CmdCompareChanged(newMap[cmdName], oldMap[cmdName]) {
			diff[cmdName] = Updated
		} else {
			diff[cmdName] = Equal
		}
	}

	return diff
}

func CmdCompareChanged(opt1, opt2 *discordgo.ApplicationCommand) bool {
	if opt1.Description != opt2.Description {
		log.Println("Description changed from ", opt2.Description, " to ", opt1.Description)
		return true
	}

	if opt1.Options != nil && opt2.Options != nil {
		if len(opt1.Options) != len(opt2.Options) {
			log.Println("Option size changed!")
			return true
		} else {
			for i := range opt1.Options {
				if OptionsComparedChanged(opt1.Options[i], opt2.Options[i]) {
					return true
				}
			}
		}
	}

	return false
}

func OptionsComparedChanged(opt1, opt2 *discordgo.ApplicationCommandOption) bool {
	if len(opt1.Options) != len(opt2.Options) {
		return true
	} else if opt1.Name != opt2.Name {
		log.Println("Name changed from ", opt2.Name, " to ", opt1.Name)
		return true
	} else if opt1.Description != opt2.Description {
		log.Println("Descripton changed from ", opt2.Description, " to ", opt1.Description)
		return true
	} else if opt1.Type != opt2.Type {
		log.Println("Type changed from ", opt2.Type, " to ", opt1.Type)
		return true
	} else if opt1.Required != opt2.Required {
		log.Println("Required changed from ", opt2.Required, " to ", opt1.Required)
		return true
	} else if opt1.Autocomplete != opt2.Autocomplete {
		log.Println("Autocomplete changed from ", opt2.Autocomplete, " to ", opt1.Autocomplete)
		return true
	} else if opt1.Options != nil && opt2.Options != nil {
		if len(opt1.Options) != len(opt2.Options) {
			return true
		} else {
			for i := range opt1.Options {
				if OptionsComparedChanged(opt1.Options[i], opt2.Options[i]) {
					return true
				}
			}
		}
	}
	return false
}
