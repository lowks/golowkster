package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

func main() {
  var filename string
  app := cli.NewApp()

  app.Name = "greet"
  app.Usage = "fight the loneliness!"
  app.Version = "0.0.1"

  app.Flags = []cli.Flag {
    cli.StringFlag {
	Name: "filename",
	Value: "greet.go",
	Usage: "Filename `FILE`",
	Destination: &filename,
    },
  }

  app.Action = func(c *cli.Context) error {
    if filename == "greet.go" {
    	fmt.Println("Hello friend!")
    } else {
	fmt.Println("On no. No more friends")
    }
    return nil
  }

  app.Run(os.Args)
}
