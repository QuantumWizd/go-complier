package chat

import (
	"fmt"
	"strings"
)

const (
	TypeMsg  = "MSG"
	TypeSys  = "SYS"
	TypeHist = "HIST"
	TypeWho  = "WHO"
)

type Message struct {
	Type string
	From string
	Body string
}

func Format(m Message) string {
	return fmt.Sprintf("%s:%s:%s\n", m.Type, m.From, m.Body)
}

func Parse(line string) (Message, error) {
	line = strings.TrimSpace(line)

	if line == "" {
		return Message{}, fmt.Errorf("empty message")
	}

	parts := strings.SplitN(line, ":", 3)

	if len(parts) != 3 {
		return Message{}, fmt.Errorf("invalid message format: %s", line)
	}

	msgType := strings.TrimSpace(parts[0])
	from    := strings.TrimSpace(parts[1])
	body    := strings.TrimSpace(parts[2])

	if msgType != TypeMsg && msgType != TypeSys && msgType != TypeHist && msgType != TypeWho {
		return Message{}, fmt.Errorf("unknown message type: %s", msgType)
	}

	if from == "" {
		return Message{}, fmt.Errorf("sender name is empty")
	}

	return Message{
		Type: msgType,
		From: from,
		Body: body,
	}, nil
}