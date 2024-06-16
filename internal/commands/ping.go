package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func init() {
	PingCommand := Command{
		Name:       "ping",
		Aliases:    []string{},
		FuncToCall: Ping,
		ArgsParser: PingParser,
	}
	CommandsList = append(CommandsList, PingCommand)
}

func PingParser(args []string) []string {
	return nil
}

func Ping(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Ping (%d ms)", session.HeartbeatLatency()/1000000))
	return 
}
