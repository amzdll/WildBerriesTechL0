package pg

import "errors"

var ErrNotFound = errors.New("record not found")
var ErrDuplicateKey = errors.New("duplicate key error")
