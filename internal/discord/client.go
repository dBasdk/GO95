package discord

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Client struct {
	Session *discordgo.Session
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("token is empty, please provide a valid token in config.yml")
	}
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return &Client{Session: session}, nil
}

func (c *Client) Open() error {
	err := c.Session.Open()
	if err != nil {
		return err
	}
	log.Println("Bot is connected")
	return nil
}

func (c *Client) AddIntent(intent []discordgo.Intent) {
	for _, i := range intent {
		c.Session.Identify.Intents |= i
	}
}

func (c *Client) AddHandler(handler interface{}) {
	c.Session.AddHandler(handler)
}

func (c *Client) Close() error {
	err := c.Session.Close()
	if err != nil {
		return err
	}
	log.Println("Bot is disconnected")
	return nil
}
