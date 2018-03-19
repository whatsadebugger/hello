package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "add",
			Usage:  "add two arguments",
			Action: add,
		},
		cli.Command{
			Name:   "multiply",
			Usage:  "multiply two arguments",
			Action: multiply,
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

func add(c *cli.Context) error {
	args := c.Args()
	a, _ := strconv.Atoi(args[0])

	b, _ := strconv.Atoi(args[1])
	fmt.Println(a + b)

	return nil
}

func multiply(c *cli.Context) error {
	args := c.Args()
	a, _ := strconv.Atoi(args[0])

	b, _ := strconv.Atoi(args[1])
	fmt.Println(a * b)

	return nil
}
