package main

import (
	"fmt"
	"golang-structure-started/mypackage"

	"github.com/spf13/cobra"
)

func main() {

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello")

			mypackage.PrintSomething()
		},
	}

	fmt.Println("Calling cmd.Execute()")
	cmd.Execute()
}
