package commands

import "fmt"

func init() {
	CampfireCommand := Command{
		Name:       "campfire",
		Aliases:    []string{"camp", "fire"},
		FuncToCall: Ping,
	}
	CommandsList = append(CommandsList, CampfireCommand)
}

func Campfire() {
	fmt.Println("Campfire")
}