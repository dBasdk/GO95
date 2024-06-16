package commands

type Command struct {
	Name       string
	Aliases    []string
	FuncToCall func()
}

var CommandsList []Command
