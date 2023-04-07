package ncdeck

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Boards []Board

type Board struct {
	client      *Client
	Title       string        `json:"title"`
	Owner       interface{}   `json:"owner"`
	Color       string        `json:"color"`
	Archived    bool          `json:"archived"`
	Labels      Labels        `json:"labels"`
	ACL         []interface{} `json:"acl"`
	Permissions struct {
		PermissionRead   bool `json:"PERMISSION_READ"`
		PermissionEdit   bool `json:"PERMISSION_EDIT"`
		PermissionManage bool `json:"PERMISSION_MANAGE"`
		PermissionShare  bool `json:"PERMISSION_SHARE"`
	} `json:"permissions"`
	Users        []interface{} `json:"users"`
	Shared       int           `json:"shared"`
	Stacks       Stacks        `json:"stacks"`
	DeletedAt    int           `json:"deletedAt"`
	LastModified interface{}   `json:"lastModified"`
	Settings     interface{}   `json:"settings"`
	ID           int           `json:"id"`
	ETag         string        `json:"ETag"`
}

// Create card
func (b Board) CreateCard(stackid int, title, description string, order int, duedate string) (Card, error) {

	var stack = Stack{ID: stackid, client: b.client, BoardID: b.ID}

	card, err := stack.CreateCard(title, description, order, duedate)
	if err != nil {
		return Card{}, err
	}

	return card, err
}

// Create a label
func (b Board) CreateLabel(title, color string) (Label, error) {

	url := b.client.BaseURL + fmt.Sprintf("/boards/%v/labels", b.ID)
	var reqBody = []byte(fmt.Sprintf(`{"title": "%v", "color": "%v"}`, title, color))

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return Label{}, err
	}

	req.Body = ioutil.NopCloser(strings.NewReader(string(reqBody)))

	var label Label
	err = b.client.do(req, &label)
	if err != nil {
		return Label{}, err
	}

	label.client = b.client

	b.Labels = append(b.Labels, label)

	return label, nil
}

// Create a Stack
func (b Board) CreateStack(title string, order int) (Stack, error) {

	url := b.client.BaseURL + fmt.Sprintf("/boards/%v/stacks", b.ID)
	var reqBody = []byte(fmt.Sprintf(`{"title":"%v", "order": "%v"}`, title, order))

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return Stack{}, err
	}

	req.Body = ioutil.NopCloser(strings.NewReader(string(reqBody)))

	var stack Stack
	err = b.client.do(req, &stack)
	if err != nil {
		return Stack{}, err
	}

	stack.client = b.client

	b.Stacks = append(b.Stacks, stack)

	return stack, nil
}

// Get all stacks of this board
func (b *Board) GetStacks() (Stacks, error) {

	url := b.client.BaseURL + fmt.Sprintf("/boards/%v/stacks", b.ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Stacks{}, err
	}

	err = b.client.do(req, &b.Stacks)
	if err != nil {
		return Stacks{}, err
	}

	for _, stack := range b.Stacks {
		stack.client = b.client
	}

	return b.Stacks, nil
}

// Get archived stacks
func (b *Board) GetStacksArchived() error {

	url := b.client.BaseURL + fmt.Sprintf("/boards/%v/stacks/archived", b.ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	err = b.client.do(req, &b.Stacks)
	if err != nil {
		return err
	}

	for _, stack := range b.Stacks {
		stack.client = b.client
	}

	return nil
}

// Get information about this board
// useful if a board was updated by UI
func (b *Board) Get() error {

	url := fmt.Sprintf("%v/boards/%v", b.client.BaseURL, b.ID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	return b.client.do(req, &b)
}

// Delete this board
func (b *Board) Delete() error {

	url := fmt.Sprintf("%v/boards/%v", b.client.BaseURL, b.ID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return b.client.do(req, &b)
}

// Undo board deletion
func (b *Board) UndoDelete() error {

	url := fmt.Sprintf("%v/boards/%v/undo_delete", b.client.BaseURL, b.ID)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	return b.client.do(req, &b)
}

// Update informations about this board
func (b *Board) Update() error {

	url := fmt.Sprintf("%v/boards/%v", b.client.BaseURL, b.ID)
	var reqBody = []byte(fmt.Sprintf(`{"title":"%v","color": "%v","archived": %v}`, b.Title, b.Color, b.Archived))

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	req.Body = ioutil.NopCloser(strings.NewReader(string(reqBody)))

	return b.client.do(req, &b)
}
