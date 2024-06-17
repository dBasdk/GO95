package discord

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

func NewSession(token string) (*discordgo.Session, error) {
	if token == "" {
		return nil, errors.New("token is empty, please provide a valid token in config.yml")
	}
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return session, nil
}


func AddIntent(session *discordgo.Session, intent []discordgo.Intent) {
	for _, i := range intent {
		session.Identify.Intents |= i
	}
}

