package commands

import(
	"github.com/bwmarrin/discordgo"
	"fmt"
)

func init() {
	CampfireCommand := Command{
		Name:       "campfire",
		Aliases:    []string{"camp", "fire"},
		FuncToCall: Campfire,
		ArgsParser: CampfireParser,
	}
	CommandsList = append(CommandsList, CampfireCommand)
}

func CampfireParser (args []string) []string {
	return args
}

func Campfire(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	fmt.Println("Campfire")
}
