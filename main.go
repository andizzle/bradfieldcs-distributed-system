package main

import (
	"fmt"

	"github.com/andizzle/kv/cmd"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Prompt{
		Label: "bdstore",
	}

	for {
		input, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		cmd.Exec(input)
	}
}
