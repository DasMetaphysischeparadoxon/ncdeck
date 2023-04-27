# Nextcloud Deck API 

[![Go Report Card](https://goreportcard.com/badge/github.com/DasMetaphysischeparadoxon/ncdeck)](https://goreportcard.com/report/github.com/DasMetaphysischeparadoxon/ncdeck)

A Go package to interact with [Nextcloud Deck REST API](https://deck.readthedocs.io/en/latest/API/) for Nextcloud [Deck](https://github.com/nextcloud/deck) 

> this project is still in progress, please use with caution!

## Installation

```bash
go get github.com/DasMetaphysischeparadoxon/ncdeck
```

## Example

```go
package main

import (
	"fmt"

	"github.com/DasMetaphysischeparadoxon/ncdeck"
)

func main() {

	deck := ncdeck.NewClient("MyUser", "MyPassword", "https://my.nextcloud.com")

	fmt.Println("[*] get boards")
	boards, err := deck.GetBoards()
	if err != nil {
		fmt.Println("Can't get boards:", err)
	}

	for _, board := range boards {
		fmt.Printf("\t- %v\n", board.Title)
	}

	fmt.Println("[*] create board")
	board, err := deck.CreateBoard("MyBoard", "0099cc")
	if err != nil {
		fmt.Println("Can't create board:", err)
	}

	fmt.Println("[*] create stack")
	stack, err := board.CreateStack("TODO", 1)
	if err != nil {
		fmt.Println("Can't create stack:", err)
	}

	fmt.Println("[*] create card")
	card, err := stack.CreateCard("My Title", "My Description", 1, "")
	if err != nil {
		fmt.Println("Can't create card:", err)
	}

	fmt.Println("[*] update card")
	card.Title = "New Title"
	err = card.Update()
	if err != nil {
		fmt.Println("Can't update card:", err)
	}
}
```

For more examples please look into the test file ```cmd/test_client/main.go```.

## Documentation

Please use ```go doc``` be like:

```bash
go doc ncdeck.client
```
