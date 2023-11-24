package ncdeck

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Cards []Card

type Card struct {
	boardID         int
	client          *Client
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	StackID         int         `json:"stackId"`
	Type            string      `json:"type"`
	LastModified    int         `json:"lastModified"`
	CreatedAt       int         `json:"createdAt"`
	Labels          Labels      `json:"labels"`
	AssignedUsers   interface{} `json:"assignedUsers"`
	Attachments     interface{} `json:"attachments"`
	AttachmentCount interface{} `json:"attachmentCount"`
	Owner           interface{} `json:"owner"`
	Order           int         `json:"order"`
	Archived        bool        `json:"archived"`
	Duedate         string      `json:"duedate"`
	DeletedAt       int         `json:"deletedAt"`
	CommentsUnread  int         `json:"commentsUnread"`
	ID              int         `json:"id"`
	Overdue         int         `json:"overdue"`
}

// Get informations about this card
// useful if this card was updated by UI
func (c *Card) Get() error {

	var url = c.client.BaseURL + fmt.Sprintf("/boards/%v/stacks/%v/cards/%v", c.boardID, c.StackID, c.ID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	return c.client.do(req, &c)
}

// Update informations about this card
func (c *Card) Update() error {

	var url = c.client.BaseURL + fmt.Sprintf("/boards/%v/stacks/%v/cards/%v", c.boardID, c.StackID, c.ID)

	if c.Duedate == "" {
		// If duedate is empty, the update will failed
		c.Duedate = "null"
	}

	var reqBody = fmt.Sprintf(`{"title":"%v","description":%#v,"type":"plain","order":%v,"duedate":%v,"owner":%#v}`, c.Title, c.Description, c.Order, c.Duedate, c.Owner)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	req.Body = io.NopCloser(strings.NewReader(reqBody))

	return c.client.do(req, &c)
}

// Delete this card
func (c *Card) Delete() error {

	var url = c.client.BaseURL + fmt.Sprintf("/boards/%v/stacks/%v/cards/%v", c.boardID, c.StackID, c.ID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.client.do(req, &c)
}
