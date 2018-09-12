package message

import (
	"fmt"
)

type Prefixer struct {
	prefix   string
	message Messager
}

func Prefix(prefix string, message Messager) Messager {
	return Prefixer{
		prefix: prefix,
		message: message,
	}
}

var _ Messager = &Prefixer{}

func (m Prefixer) Message() (string, error) {
	msg, err := m.message.Message()
	if err != nil {
		return "", err
	}

	if msg == "" {
		return "", nil
	}

	return fmt.Sprintf("%s%s", m.prefix, msg), nil
}
