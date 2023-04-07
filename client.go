// a go package to interact with Nextcloud Deck API
package ncdeck

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	BaseURL   string
	username  string
	password  string
	Client    *http.Client
	debugMode bool
	ctx       context.Context
}

// Create new Nextcloud Deck Client
func NewClient(username, password, baseURL string) *Client {
	return &Client{
		BaseURL:  baseURL + "/index.php/apps/deck/api/v1.2",
		username: username,
		password: password,
		Client:   &http.Client{Timeout: time.Second * 60},
	}
}

// Execute HTTP requests
func (c Client) do(req *http.Request, obj interface{}) error {

	req.Header.Add("OCS-APIRequest", "true")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("status code: %d\nbody: %v", res.StatusCode, res.Body)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// log.Printf("URL: %v\nResponse body: %v\n", req.URL, string(body))

	return json.Unmarshal(body, &obj)
}

// Get Stacks of Board ID
func (c Client) GetStacks(boardid, stackid int) (Stacks, error) {

	board := Board{ID: boardid}
	board.Stacks = append(board.Stacks, Stack{ID: stackid, BoardID: boardid, client: &c})

	return board.GetStacks()
}

// Get cards from a stack
func (c Client) GetCards(boardid, stackid int) (Cards, error) {

	board := Board{ID: boardid}
	board.Stacks = append(board.Stacks, Stack{ID: stackid, BoardID: boardid, client: &c})

	return board.Stacks[0].GetCards()
}

// Create card on specific board and stack
func (c Client) CreateCard(boardid, stackid int, title, description string, order int, duedate string) (Card, error) {

	board := Board{ID: boardid}
	board.Stacks = append(board.Stacks, Stack{ID: stackid, BoardID: boardid, client: &c})

	card, err := board.Stacks[0].CreateCard(title, description, order, duedate)
	if err != nil {
		return Card{}, err
	}

	return card, err
}

// Create a new board
func (c Client) CreateBoard(title, color string) (Board, error) {

	url := c.BaseURL + "/boards"
	body := fmt.Sprintf(`{"title":"%v", "color": "%v"}`, title, color)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return Board{}, err
	}

	req.Body = ioutil.NopCloser(strings.NewReader(body))

	var board Board

	err = c.do(req, &board)
	if err != nil {
		return Board{}, err
	}

	board.client = &c

	return board, nil
}

// Get all boards
func (c Client) GetBoards() (Boards, error) {

	url := c.BaseURL + "/boards"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Boards{}, err
	}

	var boards Boards

	err = c.do(req, &boards)
	if err != nil {
		return Boards{}, err
	}

	for _, board := range boards {
		board.client = &c
	}

	return boards, nil
}

// Get board by id
func (c Client) GetBoard(boardid int) (Board, error) {

	url := c.BaseURL + fmt.Sprintf("/boards/%v", boardid)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Board{}, err
	}

	var board Board

	board.client = &c

	err = c.do(req, &board)
	if err != nil {
		return Board{}, err
	}

	return board, nil
}
