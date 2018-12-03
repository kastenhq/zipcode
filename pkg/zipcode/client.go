package zipcode

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const contentType = "text/html"

type client struct {
	baseURL url.URL
}

func newClient(host string) *client {
	return &client{
		baseURL: url.URL{
			Scheme: "http",
			Host:   "localhost:8000",
		},
	}
}

func (c *client) resetOrders() error {
	u := c.baseURL
	u.Path = "reset"
	_, err := http.Post(u.String(), contentType, nil)
	return err
}

func (c *client) insertOrder(zips ...string) error {
	u := c.baseURL
	u.Path = "orders"
	u.RawQuery = url.Values{"zip": zips}.Encode()
	r, err := http.Post(u.String(), contentType, nil)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	log.Println("Insert Order Response", string(b))
	return nil
}
