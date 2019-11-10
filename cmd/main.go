package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mo-work/go-technical-test-for-claudia/image"
	"github.com/mo-work/go-technical-test-for-claudia/input"
)

func main() {
	in := input.New(bufio.NewScanner(os.Stdin))
	xAxis, yAxis, err := in.GetImageSize()
	if err != nil {
		fmt.Printf("invalid image value: %s\n", err)
	}

	bitmap := image.New(xAxis, yAxis)

	commandChan := make(chan input.Command)
	errChan := make(chan error)
	go in.GetEditActions(commandChan, errChan)

	for {
		select {
		case command := <-commandChan:
			switch command.Action {
			case "L":
				if err := bitmap.Set(command.Coords[0], command.Coords[1], command.Char); err != nil {
					fmt.Println(err)
				}
			case "V":
				if err := bitmap.SetMultiY(command.Coords[0], command.Coords[1], command.Coords[2], command.Char); err != nil {
					fmt.Println(err)
				}
			case "H":
				if err := bitmap.SetMultiX(command.Coords[0], command.Coords[1], command.Coords[2], command.Char); err != nil {
					fmt.Println(err)
				}
			case "S":
				fmt.Println(bitmap.Pretty())
			case "C":
				bitmap.Clear()
			case "":
				os.Exit(0)
			default:
				fmt.Println("invalid action")
			}

		case err := <-errChan:
			fmt.Println(err)
		}
	}
}
