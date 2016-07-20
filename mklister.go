package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/urfave/cli.v1"
)

func PrintContents(path string, level int) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() {
			indentation := ""

			for i := 0; i < level; i++ {
				indentation += " "
			}

			fmt.Println(indentation + file.Name() + "/")

			PrintContents((path + "/" + file.Name()), (level + 1))

		} else {
			indentation := " "

			for i := 0; i < level; i++ {
				indentation += " "
			}

			symPath := SymLink(file, path)

			if symPath != "" {
				fmt.Println(indentation + file.Name() + "*" + " (" + symPath + ")")
			} else {
				fmt.Println(indentation + file.Name())
			}
		}
	}

	return nil
}

func SymLink (file os.FileInfo, path string) string {
	if file.Mode() & os.ModeSymlink == os.ModeSymlink {
		symPath, err := os.Readlink(path + "/" + file.Name())

		if err != nil {
			log.Fatal(err)
		}

		return symPath
	}

	return ""
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
		fmt.Println(directory + "/")

		if recursive {
			PrintContents(directory, 1)
		} else {
			files, err := ioutil.ReadDir(directory)

			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				fmt.Println(" " + file.Name())
			}
		}

		return nil
	}

	app.Run(os.Args)
}
