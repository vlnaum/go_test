package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

type client struct {
	address string
	timeout time.Duration
	in      io.Reader
	out     io.Writer
	conn    net.Conn
}

func (c *client) Connect() (err error) {
	c.conn, err = net.DialTimeout("tcp", c.address, c.timeout)
	return
}

func (c *client) Close() (err error) {
	err = c.conn.Close()
	return
}

func (c *client) Send() (err error) {
	_, err = io.Copy(c.conn, c.in)
	return
}

func (c *client) Receive() (err error) {
	_, err = io.Copy(c.out, c.conn)
	return
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient { //nolint:interfacer
	return &client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
