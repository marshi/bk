package main

import (
	"os"
	"github.com/urfave/cli"
	"bk/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "bk"
	app.Usage = "bookmark current directory."
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name: "show",
			Usage: "",
			Action: commands.Show,
		},
		{
			Name: "save",
			Usage: "",
			Action: commands.Save,
		},
		{
			Name: "del",
			Usage: "",
			Action: commands.Delete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "path",
				},
			},
		},
	}
	if 1 < len(os.Args) {
		app.Run(os.Args)
	}
}
