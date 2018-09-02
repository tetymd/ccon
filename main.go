package main

import (
    "fmt"
    "os"
    "github.com/c-bata/go-prompt"
)

func executer(in string) {
    if in == "exit" {
        os.Exit(0)
    } else {
        fmt.Println("Your input: " + in)

    }
}

func completer(d prompt.Document) []prompt.Suggest {
    s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
    p := prompt.New (
        executer,
        completer,
        prompt.OptionPrefix(">>> "),
        prompt.OptionTitle("sql-prompt"),
    )
    p.Run()
}
