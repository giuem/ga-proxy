package ga

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

func PageView(data CommonData) error {
	data.HitType = "pageview"

	v, err := query.Values(data)
	if err != nil {
		return errors.Wrap(err, "could not encode query")
	}

	err = send(v.Encode())

	return nil
}

func Detect() error {
	err := send("")
	return err
}
