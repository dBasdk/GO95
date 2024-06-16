package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"GO95/internal/discord"
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
	dclient, err := discord.NewClient(Token)
	checkNilErr(err)
	dclient.AddHandler(events.MessageCreate)
	dclient.AddIntent([]discordgo.Intent{discordgo.IntentsGuildMessages})
	err = dclient.Open()
	checkNilErr(err)
	defer dclient.Close()
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
