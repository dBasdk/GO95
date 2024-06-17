package discord

// It should be used to answer to a command
// It should answer both type of message, Interaction and Message

import(
	"github.com/bwmarrin/discordgo"
)

type EventType int

const (
    MessageType EventType = iota
    InteractionType
)

type Event struct {
    Type         EventType
    Message      *discordgo.MessageCreate
    Interaction  *discordgo.InteractionCreate
}

type ResponseType int

const (
	Embed ResponseType = iota
	SimpleMessage
)

type ResponseData struct  {
	Type  			ResponseType
	Embed			*discordgo.MessageEmbed
	SimpleMessage	string
}

func Reply(session *discordgo.Session, event Event, data ResponseData) (e error) {
	switch event.Type {
    case MessageType:
    	switch data.Type {
    	case Embed:
    		session.ChannelMessageSendEmbedReply(
    			event.Message.ChannelID,
    			data.Embed,
    			event.Message.Reference())
    	case SimpleMessage:
    		session.ChannelMessageSendReply(
    			event.Message.ChannelID,
    			data.SimpleMessage,
    			event.Message.Reference())
    	}
    case InteractionType:
    	switch data.Type {
    	case Embed:
    		session.InteractionRespond(
    			event.Interaction.Interaction,
    			&discordgo.InteractionResponse{
    				Type: discordgo.InteractionResponseChannelMessageWithSource,
    				Data: &discordgo.InteractionResponseData{
    					Embeds: []*discordgo.MessageEmbed{
    						data.Embed},
    				},
    			},
    		)
    	case SimpleMessage:
    		session.InteractionRespond(
    			event.Interaction.Interaction,
    			&discordgo.InteractionResponse{
    				Type: discordgo.InteractionResponseChannelMessageWithSource,
    				Data: &discordgo.InteractionResponseData{
    					Content: data.SimpleMessage},
    			},
    		)
    	}
    }
    return 
}