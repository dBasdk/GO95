package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
	"strings"
)



type Config struct {
	Token string
}

type CommandInfos struct {
	Command string
	Args []string
}

var (
	Token string
	config Config
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


func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	var parts string
	if strings.HasPrefix(msg.Content, "g.") {
		parts = strings.Replace(msg.Content, "g.", "", 1)
	} else {
		if strings.HasPrefix(msg.Content, "<@1251530001975869481>"){
			parts = strings.Replace(msg.Content, "<@1251530001975869481>", "", 1)
		}
	}
	
	fields := strings.Fields(parts)


	var infos CommandInfos
	if len(fields) > 0 {
		infos.Command = fields[0]
		infos.Args = fields[1:]
	} else {
		return
	}

	if msg.Author.Bot {
		return
	}
}

func main() {
	if Token == "" {
		fmt.Println("Token not found in config.yml, please provide one")
		os.Exit(1)
	}

	dg, err := discordgo.New("Bot " + Token)
	checkNilErr(err)
	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	checkNilErr(err)
	defer dg.Close()
	
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}


