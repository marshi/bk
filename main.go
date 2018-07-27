package main

import (
	"bk/commands"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bk"
	app.Usage = "bookmark current directory."
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:   "show",
			Usage:  "",
			Action: commands.Show,
		},
		{
			Name:   "save",
			Usage:  "",
			Action: commands.Save,
		},
		{
			Name:   "del",
			Usage:  "",
			Action: commands.Delete,
		},
	}
	if 1 < len(os.Args) {
		app.Run(os.Args)
	}
}
