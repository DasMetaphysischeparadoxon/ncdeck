package main

import (
	"fmt"
	"os"

	"github.com/DasMetaphysischeparadoxon/ncdeck"
)

var (
	username, password, baseurl string
)

func init() {
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
	baseurl = os.Getenv("BASEURL")
}

// ignore gocyclo because its a test program
//
//gocyclo:ignore
func main() {

	fmt.Println("[*] create client")
	ncd := ncdeck.NewClient(username, password, baseurl)

	fmt.Println("[*] get boards")
	boards, err := ncd.GetBoards()
	if err != nil {
		fmt.Println("Can't get boards:", err)
	}

	for _, board := range boards {
		fmt.Printf("\t- %v\n", board.Title)
	}

	fmt.Println("[*] create board")
	board, err := ncd.CreateBoard("Test123", "fff000")
	if err != nil {
		fmt.Println("Fail to create board:", err)
	}
	fmt.Printf("\tcreated board: %v\n", board.Title)

	fmt.Println("[*] try to get board by id")
	board, err = ncd.GetBoard(board.ID)
	if err != nil {
		fmt.Println("Fail to get board:", err)
	}

	color := "fffeee"
	fmt.Printf("[*] update board color from %v to %v\n", board.Color, color)
	board.Color = color
	err = board.Update()
	if err != nil {
		fmt.Println("Fail to update board:", err)
	}

	fmt.Println("[*] get labels from board")
	for _, label := range board.Labels {
		fmt.Printf("\t Title: %v Color: %v\n", label.Title, board.Color)
	}

	fmt.Println("[*] try to create a label for board:", board.Title)
	label, err := board.CreateLabel("TestLabel", "fffff")
	if err != nil {
		fmt.Println("Fail to create label:", err)
	}
	fmt.Printf("\tTitle: %v\n\tColor: %v\n", label.Title, label.Color)

	color = "7877e6"
	fmt.Printf("[*] try to change label color from %v to %v\n", label.Color, color)
	label.Color = color
	err = label.Update()
	if err != nil {
		fmt.Println("Fail to update label:", err)
	}
	fmt.Printf("\tTitle: %v\n\tColor: %v\n", label.Title, label.Color)

	fmt.Println("[*] try to get label")
	err = label.Get()
	if err != nil {
		fmt.Println("Fail to update label:", err)
	}
	fmt.Printf("\tTitle: %v\n\tColor: %v\n", label.Title, label.Color)

	fmt.Println("[*] try to delete label:", label.Title)
	err = label.Delete()
	if err != nil {
		fmt.Println("Fail to delete label:", err)
	}
	fmt.Printf("\tTitle: %v\n\tColor: %v\n", label.Title, label.Color)

	fmt.Println("[*] create stacks for board")
	stack1, err := board.CreateStack("TODO", 999)
	if err != nil {
		fmt.Println("Fail to create stack1", err)
	}

	fmt.Printf("\tcreate stack: %v\n", stack1.Title)
	stack2, err := board.CreateStack("DINE", 999)
	if err != nil {
		fmt.Println("Fail to create stack2", err)
	}

	fmt.Printf("\tcreate stack with typo: %v\n", stack2.Title)
	stack3, err := board.CreateStack("READY TO DIE", 999)
	if err != nil {
		fmt.Println("Fail to create stack3", err)
	}
	fmt.Printf("\tcreate stack to delete: %v\n", stack3.Title)

	fmt.Println("[*] try to delete", stack3.ID)
	err = stack3.Delete()
	if err != nil {
		fmt.Println("Fail to create stack2", err)
	}

	fmt.Println("[*] get stacks from board")
	stacks, err := board.GetStacks()
	if err != nil {
		fmt.Println("Fail to get stacks from board", err)

	}

	for _, stack := range stacks {
		fmt.Printf("\t- %v\n", stack.Title)
	}

	title := "DONE"
	fmt.Printf("[*] update stack2 because typo from %v to %v\n", stack2.Title, title)
	stack2.Title = title
	err = stack2.Update()
	if err != nil {
		fmt.Println("Fail to update Stack", err)
	}

	fmt.Println("[*] create cards on stack:", stack1.Title)
	card1, err := stack1.CreateCard("Card1", "TEST", 999, "null")
	if err != nil {
		fmt.Println("Fail to create card on Stack2", err)
	}
	fmt.Printf("\tcreated card: %v\n", card1.Title)

	card2, err := stack2.CreateCard("Card2", "lorem ipsum", 999, "null")
	if err != nil {
		fmt.Println("Failed to create card on Stack2", err)
	}
	fmt.Printf("\tcreated card: %v\n", card2.Title)

	order := 10
	fmt.Printf("[*] update card order from %v to %v\n", card1.Order, order)
	card1.Order = order
	err = card1.Update()
	if err != nil {
		fmt.Println("Fail to update card", err)
	}

	fmt.Println("[*] delete card", card1.Title)
	err = card1.Delete()
	if err != nil {
		fmt.Println("Failed to delete card", err)
	}

	fmt.Println("[*] update stack")
	err = stack1.Update()
	if err != nil {
		fmt.Println("Failed to update stack", err)
	}

	fmt.Println("[*] check new cards in stack2")
	cards, err := stack2.GetCards()
	if err != nil {
		fmt.Println("Failed to get cards from stack", err)
	}

	for _, card := range cards {
		fmt.Printf("\t- %v\n\t  %v\n", card.Title, card.Description)
	}

	fmt.Println("[*] create card in", stack2.Title)
	_, err = stack2.CreateCard("Card3", "test123text", 20, "")
	if err != nil {
		fmt.Println("Failed to create card", err)
	}

	fmt.Println("[*] get stack", stack2.Title)
	err = stack2.Get()
	if err != nil {
		fmt.Println("Failed to get stack", err)
	}

	fmt.Println("[*] create multiline description in Card")
	_, err = stack2.CreateCard("Multiline Card", "Line1\nLine2", 99, "null")
	if err != nil {
		fmt.Println("Failed to create card", err)
	}

	fmt.Printf("[*] get all archived stacks\n")
	err = board.GetStacksArchived()
	if err != nil {
		fmt.Println("Failed to get all archived stacks", err)
	}

	fmt.Println("[*] delete board:", board.ID)
	err = board.Delete()
	if err != nil {
		fmt.Println("Fail to delete board", err)
	}

	fmt.Println("[*] undo delete board:", board.ID)
	err = board.UndoDelete()
	if err != nil {
		fmt.Println("Fail to undo delete board", err)
	}

	fmt.Println("[*] delete board finally:", board.ID)
	err = board.Delete()
	if err != nil {
		fmt.Println("Failed to delete board", err)
	}

}
