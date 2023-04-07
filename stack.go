package ncdeck

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Stacks []Stack

type Stack struct {
	client       *Client
	Title        string `json:"title"`
	BoardID      int    `json:"boardId"`
	DeletedAt    int    `json:"deletedAt"`
	LastModified int    `json:"lastModified"`
	Cards        Cards  `json:"cards"`
	Order        int    `json:"order"`
	ID           int    `json:"id"`
}

// Create card on this stack
func (s Stack) CreateCard(title, description string, order int, duedate string) (Card, error) {

	url := s.client.BaseURL + fmt.Sprintf("/boards/%v/stacks/%v/cards", s.BoardID, s.ID)

	var reqBody = []byte(fmt.Sprintf(`{"title":"%v", "type": "plain", "order": %v, "description": %#v}`, title, order, description))

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return Card{}, err
	}

	req.Body = ioutil.NopCloser(strings.NewReader(string(reqBody)))

	var card Card
	card.client = s.client
	card.boardID = s.BoardID
	err = s.client.do(req, &card)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}

// Update stacks informations
func (s *Stack) Update() error {

	url := s.client.BaseURL + fmt.Sprintf("/boards/%v/stacks/%v", s.BoardID, s.ID)
	var reqBody = []byte(fmt.Sprintf(`{"title":"%v", "order": "%v"}`, s.Title, s.Order))

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	req.Body = ioutil.NopCloser(strings.NewReader(string(reqBody)))

	return s.client.do(req, &s)
}

// Get all cards from this stack
func (s *Stack) GetCards() (Cards, error) {

	url := s.client.BaseURL + fmt.Sprintf("/boards/%v/stacks/%v", s.BoardID, s.ID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Cards{}, err
	}

	err = s.client.do(req, &s)
	if err != nil {
		return Cards{}, err
	}

	return s.Cards, nil
}

// Get stacks details
// useful if stack was updated via UI
func (s *Stack) Get() error {

	url := s.client.BaseURL + fmt.Sprintf("/boards/%v/stacks/%v", s.BoardID, s.ID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	return s.client.do(req, &s)
}

// Delete this stack
func (s *Stack) Delete() error {

	url := s.client.BaseURL + fmt.Sprintf("/boards/%v/stacks/%v", s.BoardID, s.ID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return s.client.do(req, &s)
}
