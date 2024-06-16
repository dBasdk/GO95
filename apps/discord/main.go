package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"protocol"
	"strings"
	"syscall"

	"GO95/internal/commands"
	"GO95/internal/discord"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

type configStruct struct {
	Token string
}

type commandInfos struct {
	command string
	args    []string
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

func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	// log
	fmt.Println(msg.Content)
	var parts string
	if strings.HasPrefix(msg.Content, "g.") {
		parts = strings.Replace(msg.Content, "g.", "", 1)
	} else {
		if strings.HasPrefix(msg.Content, "<@1251530001975869481>") {
			parts = strings.Replace(msg.Content, "<@1251530001975869481>", "", 1)
		}
	}

	fields := strings.Fields(parts)

	var infos commandInfos
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

func main() {
	if Token == "" {
		fmt.Println("Token not found in config.yml, please provide one")
		os.Exit(1)
	}
	dclient, err := discord.NewClient(Token)
	checkNilErr(err)
	dclient.AddHandler(messageCreate)
	dclient.AddIntent([]discordgo.Intent{discordgo.IntentsGuildMessages})
	err = dclient.Open()
	checkNilErr(err)
	defer dclient.Close()
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
