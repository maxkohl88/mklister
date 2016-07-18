package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	var directory string

	app := cli.NewApp()
	app.Name = "mklister"
	app.Usage = "list the contents of a provided directory"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "path, p",
			Usage: "path to `DIRECTORY`, required",
			Destination: &directory,
		},
		cli.BoolFlag{
			Name: "recursive, r",
			Usage: "when set, list files recursively. default is off",
		},
		cli.StringFlag{
			Name: "output, o",
			Usage: "json|yml|text, default `FORMAT` is text",
		},
	}

	app.Action = func(c *cli.Context) error {
		files, err := ioutil.ReadDir(directory)

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
		}

		return nil
	}

	app.Run(os.Args)
}
