package server

type Client struct {
	Id int64
}

func NewClient() *Client {
	return &Client{}
}
