package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"discord"
	"GO95/internal/events"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

type configStruct struct {
	Token string
}

var (
	Token  string
	config configStruct
)

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
		os.Exit(1)
	}
}

func init() {
	configFile, err := os.ReadFile("config.yml")
	checkNilErr(err)

	err = yaml.Unmarshal(configFile, &config)
	checkNilErr(err)

	Token = config.Token
}

func main() {
	if Token == "" {
		fmt.Println("Token not found in config.yml, please provide one")
		os.Exit(1)
	}

	session, err := discordgo.New("Bot " + Token)

	checkNilErr(err)

	session.AddHandler(events.MessageCreate)
	
	discord.AddIntent(session, []discordgo.Intent{discordgo.IntentsGuildMessages})
	
	err = session.Open()
	checkNilErr(err)
	defer session.Close()
	
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
