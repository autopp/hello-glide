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

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number",
    Long:  `All software has versions.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hello-glide v1.0")
    },
}

var helloCmd = &cobra.Command{
    Use:   "hello",
    Short: "hello someone",
    Long:  "say hello to.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hello " + args[0])
    },
}

func main() {
    cobra.OnInitialize()
    RootCmd.AddCommand(versionCmd)
    RootCmd.AddCommand(helloCmd)
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}
