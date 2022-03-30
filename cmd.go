package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func cmd() {

var requests int
var concurrent int
var url string

  app := &cli.App{
    Name: "gz",
    Usage: "Load test with ghormehsabzi",
    Version: "v0.0.1",
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name:        "url",
        Usage:       "http url for load test",
        Aliases: []string{"u"},
        Destination: &url,
        Required: true,
      },
      &cli.IntFlag{
        Name:        "requests",
	Value:       1,
        Usage:       "Number of requests",
	Aliases: []string{"r"},
        Destination: &requests,
	DefaultText: "1",
      },
      &cli.IntFlag{
        Name:        "concurrent",
        Value:       1,
        Usage:       "Number of concurrent requests",
        Aliases: []string{"c"},
        Destination: &concurrent,
        DefaultText: "1",
      },
    },
    Action: func(c *cli.Context) error {
      if c.NArg() > 0 {
	fmt.Println(c.Args().Get(0))
      }
      ddos(url, concurrent, requests)
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
