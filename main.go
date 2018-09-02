package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/c-bata/go-prompt"
)

func executer(in string) {
    if in == "exit" {
        os.Exit(0)
    } else if in == "help" {
        fmt.Println("command format:\n    cloud [service] [action]\n")
    }

    switch parse(in)[0] {
    case "gcp":
        fmt.Println("Google Cloud Platform")
    case "aws":
        fmt.Println("Amazon Web Service")
    }
}

func completer(d prompt.Document) []prompt.Suggest {
    s := []prompt.Suggest{
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func parse(in string) []string {
    s := strings.Split(in, " ")
    return s
}

func main() {
    p := prompt.New(
        executer,
        completer,
        prompt.OptionPrefix(">>> "),
        prompt.OptionTitle("cloud console"),
    )
    p.Run()
}
