package util

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	roomId     string
	user       string
}

func (c *Client) SetUser(user string) {
	c.user = user
}

func (c *Client) GetUser() string {
	return c.user
}

func (c *Client) GetRoomId() string {
	return c.roomId
}

func (c *Client) SetRoomId(roomId string) {
	c.roomId = roomId
}

func NewClient(urlstr string) (*Client, error) {
	client := new(Client)
	parsedURL, err := url.ParseRequestURI(urlstr)
	if err != nil {
		return nil, err
	}
	client.URL = parsedURL
	client.HTTPClient = http.DefaultClient
	return client, nil
}

func (c *Client) CreateRoom(p1, p2 string) (string, error) {
	ctx := context.Background()
	var request struct {
		P1 string `json:"p1"`
		P2 string `json:"p2"`
	}
	request.P1 = p1
	request.P2 = p2
	requestByte, _ := json.Marshal(request)
	requestReader := bytes.NewReader(requestByte)
	nr, err := c.newRequest(ctx, "POST", "/api/room", requestReader)
	if err != nil {
		return "", err
	}
	res, err := c.HTTPClient.Do(nr)
	if err != nil {
		return "", err
	}
	var roomId struct {
		Val string `json:"room_id"`
	}
	if err := decodeBody(res, &roomId); err != nil {
		return "", err
	}
	c.roomId = roomId.Val
	c.user = p1
	return roomId.Val, nil
}

func (c *Client) GetRoom() (*Game, error) {
	ctx := context.Background()
	nr, err := c.newRequest(ctx, "GET", "/api/room/"+c.roomId, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(nr)
	if err != nil {
		return nil, err
	}
	game := new(Game)
	if err := decodeBody(res, &game); err != nil {
		return nil, err
	}
	return game, nil
}

func (c *Client) PutPiece(id, x, y int) (*Result, error) {
	var req struct {
		Name    string `json:"name"`
		PieceID int    `json:"piece_id"`
		PieceX  int    `json:"piece_x"`
		PieceY  int    `json:"piece_y"`
	}
	req.Name = c.GetUser()
	req.PieceID = id
	req.PieceX = x
	req.PieceY = y
	ctx := context.Background()
	nr, err := c.newRequest(ctx, "POST", "/api/room/"+c.roomId, encodeBody(req))
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(nr)
	if err != nil {
		return nil, err
	}
	result := new(Result)
	if err := decodeBody(res, &result); err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(result.Message)
	}
	return result, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func encodeBody(out interface{}) io.Reader {
	requestByte, _ := json.Marshal(out)
	return bytes.NewReader(requestByte)
}
