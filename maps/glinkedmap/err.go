package glinkedmap

import "github.com/pkg/errors"

var (
	ErrDuplicatedKey = errors.New("duplicated key in Map")
	Skip             = errors.New("skip")
)
