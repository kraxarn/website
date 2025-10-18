package mumble

import (
	"github.com/kraxarn/website/config"
	"net"
	"time"
)

type Api struct {
	conn net.Conn
	url  string
}

func NewApi() (Api, error) {
	url, err := config.MumbleUrl()
	if err != nil {
		return Api{}, err
	}

	return Api{
		conn: nil,
		url:  url,
	}, nil
}

func (a Api) Dial() error {
	timeout, err := time.ParseDuration("100ms")
	if err != nil {
		return err
	}

	a.conn, err = net.DialTimeout("tcp", a.url, timeout)
	return err
}

func (a Api) Close() error {
	if a.conn != nil {
		return a.conn.Close()
	}
	return nil
}
