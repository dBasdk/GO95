package events

import (
	"fmt"
	"protocol"
	"strings"

	"commands"

	"github.com/bwmarrin/discordgo"
)

type CommandInfos struct {
	command string
	args    []string
}

func MessageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	var parts string
	if strings.HasPrefix(msg.Content, "g.") {
		parts = strings.Replace(msg.Content, "g.", "", 1)
	} else {
		if strings.HasPrefix(msg.Content, "<@1251530001975869481>") {
			parts = strings.Replace(msg.Content, "<@1251530001975869481>", "", 1)
		}
	}

	fields := strings.Fields(parts)

	var infos CommandInfos
	if len(fields) > 0 {
		infos.command = strings.ToLower(fields[0])
		infos.args = fields[1:]
	} else {
		return
	}

	var maxValue float64 = 0
	var maxCommand commands.Command
	var curr float64
	for _, c := range commands.CommandsList {
		curr = protocol.CheckSimilarity(infos.command, c.Name)
		fmt.Println(curr, c.Name)
		if curr > maxValue {
			maxValue = curr
			maxCommand = c
			continue
		}
		for _, alias := range c.Aliases {
			curr = protocol.CheckSimilarity(infos.command, alias)
			if curr > maxValue {
				maxValue = curr
				maxCommand = c
				break
			}
		}
	}
	fmt.Println(maxCommand, maxValue)

	if msg.Author.Bot {
		return
	}
}
