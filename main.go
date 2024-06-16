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
	"protocol"
	"commands"
)



type configStruct struct {
	Token string
}

type commandInfos struct {
	command string
	args []string
}

var (
	token string
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

	token = config.Token
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
		if curr > maxValue{
			maxValue = curr
			maxCommand = c
			continue
		}
		for _, alias := range c.Aliases{
			curr = protocol.CheckSimilarity(infos.command, alias)
			if curr > maxValue{
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
	if token == "" {
		fmt.Println("Token not found in config.yml, please provide one")
		os.Exit(1)
	}

	fmt.Println(commands.CommandsList)

	dg, err := discordgo.New("Bot " + token)
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


