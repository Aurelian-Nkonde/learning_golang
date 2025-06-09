package client

type Client struct {
	name, business string
	price          int
}

func NewClient(name, business string, price int) *Client {
	return &Client{name, business, price}
}

func (cliet *Client) GetName() string {
	return cliet.name
}

func (client *Client) GetBusiness() string {
	return client.business
}

func (client *Client) GetPrice() int {
	return client.price
}

func (client *Client) GetClient() *Client {
	return client
}
