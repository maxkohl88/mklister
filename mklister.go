package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/urfave/cli.v1"
)

func PrintContents(path string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if err != nil {
			log.Fatal(err)
		}

		if file.IsDir() {
			fmt.Println(file.Name())
			PrintContents(path + "/" + file.Name())
		} else {
			fmt.Println(file.Name())
		}
	}

	return nil
}

func main() {
	var directory string
	var recursive bool

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
			Destination: &recursive,
		},
		cli.StringFlag{
			Name: "output, o",
			Usage: "json|yml|text, default `FORMAT` is text",
		},
	}

	app.Action = func(c *cli.Context) error {
		if recursive {
			PrintContents(directory)
		} else {
			files, err := ioutil.ReadDir(directory)

			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				fmt.Println(file.Name())
			}
		}

		return nil
	}

	app.Run(os.Args)
}
