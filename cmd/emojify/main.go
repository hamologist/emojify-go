package main

import (
	"bufio"
	"fmt"
	"github.com/hamologist/emojify/pkg/model"
	"github.com/hamologist/emojify/pkg/transformer"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, err := ioutil.ReadAll(reader)
	app := &cli.App{
		Name: "emojify",
		Usage: "emojifies text read from stdio",
		Action: func(c *cli.Context) error {
			resp, err := transformer.EmojifyTransformer(
				model.EmojifyPayload{
					Message: strings.TrimSpace(string(data)),
				},
			)
			if err != nil {
				return err
			}

			fmt.Printf("%s\n", resp.Message)
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
