package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "log"
  "encoding/json"
  "time"

  "gopkg.in/urfave/cli.v1"
  "github.com/ghodss/yaml"
)

type AsJSON struct {
  Name string
  Size int64
  ModifiedTime time.Time
  IsLink bool
  IsDir bool
  LinksTo string
  Children []AsJSON
}

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

func PrepareJSONNonRecursive(path string) []AsJSON {
  files, err := ioutil.ReadDir(path)

  if err != nil {
    log.Fatal(err)
  }

  list := []AsJSON{}

  for _, file := range files {
    formatted := AsJSON{
      Name: file.Name(),
      Size: file.Size(),
      ModifiedTime: file.ModTime(),
      IsLink: (file.Mode() & os.ModeSymlink == os.ModeSymlink),
      IsDir: file.IsDir(),
      LinksTo: SymLink(file, path),
    }
    list = append(list, formatted)
  }

  return list
}

func PrepareJSON(path string) []AsJSON {
  files, err := ioutil.ReadDir(path)

  if err != nil {
    log.Fatal(err)
  }

  list := []AsJSON{}

  for _, file := range files {
    formatted := AsJSON{}

    if file.IsDir() {
      formatted = AsJSON{
          Name: file.Name(),
          Size: file.Size(),
          ModifiedTime: file.ModTime(),
          IsLink: (file.Mode() & os.ModeSymlink == os.ModeSymlink),
          IsDir: file.IsDir(),
          LinksTo: SymLink(file, path),
          Children: PrepareJSON(path + "/" + file.Name()),
      }
    } else {
      formatted = AsJSON{
          Name: file.Name(),
          Size: file.Size(),
          ModifiedTime: file.ModTime(),
          IsLink: (file.Mode() & os.ModeSymlink == os.ModeSymlink),
          IsDir: file.IsDir(),
          LinksTo: SymLink(file, path),
        }
    }
    list = append(list, formatted)
  }

  return list
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
  var output string

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
      Value: "text",
      Destination: &output,
    },
  }

  app.Action = func(c *cli.Context) error {
    if output == "json" {
      formattedOutput := []AsJSON{}

      if recursive {
        formattedOutput = PrepareJSON(directory)
      } else {
        formattedOutput = PrepareJSONNonRecursive(directory)
      }

      marshalled, _ := json.MarshalIndent(formattedOutput, "", " ")

      fmt.Println(string(marshalled))

    } else if output == "yaml" {
      formattedOutput := []AsJSON{}

      if recursive {
        formattedOutput = PrepareJSON(directory)
      } else {
        formattedOutput = PrepareJSONNonRecursive(directory)
      }

      marshalledYAML, err := yaml.Marshal(formattedOutput)

      if err != nil {
        log.Fatal(err)
      }

      fmt.Println(string(marshalledYAML))

    } else {
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
    }

    return nil
  }

  app.Run(os.Args)
}
