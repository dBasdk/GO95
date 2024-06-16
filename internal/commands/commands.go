package commands

import(
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name       string
	Aliases    []string
	FuncToCall func(*discordgo.Session, *discordgo.MessageCreate, []string) 
	ArgsParser func([]string) []string
}

var CommandsList []Command
