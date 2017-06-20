package main

import (
  "github.com/spf13/cobra"
  "fmt"
  "os"
)

var RootCmd = &cobra.Command{
    Use:   "go-keisan",
    Short: "This tool is pretty cool.",
    Long:  "This tool is a great convenience.",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}
