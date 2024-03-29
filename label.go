package ncdeck

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Labels []Label

type Label struct {
	client       *Client
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Color        string `json:"color"`
	BoardID      int    `json:"boardId"`
	CardID       int    `json:"cardId"`
	LastModified int    `json:"lastModified"`
	ETag         string `json:"ETag"`
}

// Get labels from board
func (l *Label) Get() error {

	var url = l.client.BaseURL + fmt.Sprintf("/boards/%v/labels/%v", l.BoardID, l.ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	err = l.client.do(req, &l)
	if err != nil {
		return err
	}

	return nil
}

// Update label informations
func (l *Label) Update() error {

	var (
		url = l.client.BaseURL + fmt.Sprintf("/boards/%v/labels/%v", l.BoardID, l.ID)

		reqBody = fmt.Sprintf(`{"title": "%v", "color": "%v"}`, l.Title, l.Color)
	)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	req.Body = io.NopCloser(strings.NewReader(reqBody))

	return l.client.do(req, &l)
}

// Delete a label
func (l *Label) Delete() error {

	var url = l.client.BaseURL + fmt.Sprintf("/boards/%v/labels/%v", l.BoardID, l.ID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return l.client.do(req, &l)
}
