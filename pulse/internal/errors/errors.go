package errors

import (
	err "errors"
)

var (
	ErrNoMessagesFound        = err.New("no messages found")
	ErrUserNotFoundInProducer = err.New("user not found")
)
