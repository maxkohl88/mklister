package main

import (
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mklister"
	app.Usage = "list the contents of a provided directory"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "path=, p=",
			Usage: "path to `DIRECTORY`, required",
		},
	}

	app.Run(os.Args)
}
