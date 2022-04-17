package redisweb

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/rueian/rueidis"
)

func despeckle(s string) ([]string, error) {
	qstring, err := url.PathUnescape(s)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse query string")
	}
	return strings.Split(qstring, " "), err
}

func Execute(cmdparts []string) (string, error) {
	c, _ := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	defer c.Close()

	ctx := context.Background()
	cmd := c.B().Arbitrary(cmdparts...).Build()
	resp := c.Do(ctx, cmd)
	return resp.ToString()
}
