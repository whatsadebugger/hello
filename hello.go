package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "file, f",
			Usage: "file input path",
		},
	}

	app.Action = dostuff

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func dostuff(c *cli.Context) error {
	name := "someone"
	if c.NArg() > 0 && c.IsSet("file") {
		name = c.Args()[0]
	}
	fmt.Print(name, "\n")

	return nil
}
