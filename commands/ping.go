package commands

import "fmt"

func init() {
	PingCommand := Command{
		Name:       "ping",
		Aliases:    []string{""},
		FuncToCall: Ping,
	}
	CommandsList = append(CommandsList, PingCommand)
}

func Ping() {
	fmt.Println("ping executed")
}