module github.com/whatsadebugger/hello

go 1.15

require (
	github.com/urfave/cli v1.22.5
	github.com/whatsadebugger/greetings v0.0.0-00010101000000-000000000000
)

replace github.com/whatsadebugger/greetings => ../greetings
