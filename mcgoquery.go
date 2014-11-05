package mcgoquery

import (
	"github.com/lukevers/mc/mcquery"
	"strconv"
)

type Client struct {
	Host    string
	Port    int
	conn    *mcquery.Connection
}

// Create a Client and return it, connecting also.
func Create(host string, port int) (c *Client, err error) {
	// Create our client to return
	c = &Client{
		Host:    host,
		Port:    port,
		conn:    nil,
	}

	// Try and connect, and then return whatever we get back.
	c.conn, err = mcquery.Connect(host+":"+strconv.Itoa(port))
	return
}

// Basic Stat
func (c *Client) Basic() (*mcquery.Stat, error) {
	return c.conn.BasicStat()
}

// Full Stat
func (c *Client) Full() (*mcquery.Stat, error) {
	return c.conn.FullStat()
}
