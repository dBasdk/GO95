package commands

import (
	"fmt"

	"discord"

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
	// session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Ping (%d ms)", session.HeartbeatLatency()/1000000))
	content := fmt.Sprintf("Ping (%d ms)", session.HeartbeatLatency()/1000000)
	event := discord.Event{
		discord.MessageType, 
		message, 
		nil,
	}
	
	data := discord.ResponseData{
		discord.SimpleMessage,
		nil,
		content,
	}

	discord.Reply(session, event, data)
	return 
}


