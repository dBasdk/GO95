module GO95

go 1.22.4

replace protocol => ./protocol

replace commands => ./internal/commands

require (
	commands v0.0.0-00010101000000-000000000000
	github.com/bwmarrin/discordgo v0.28.1
	gopkg.in/yaml.v2 v2.4.0
	protocol v0.0.0-00010101000000-000000000000
)

require (
	github.com/adrg/strutil v0.3.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
)
